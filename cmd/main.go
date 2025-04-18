package main

import (
	"context"
	"os"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/config"
	"github.com/bryanArroyave/eventsplit/back/user-service/infra/migrations"
	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth"
	userports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/ports"
	adapters "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/infra/adapters/postgres"
	authhandlers "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/infra/handlers"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category"
	categoryports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
	categoryadapters "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/adapters/postgres"
	categoryhandlers "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/handlers"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/subscribers"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/auth/adapters/events"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/auth/middlewares"
	"github.com/bryanArroyave/golang-utils/app"
	appdtos "github.com/bryanArroyave/golang-utils/app/dtos"
	eventbroker "github.com/bryanArroyave/golang-utils/events/adapter/eventBroker"
	routerbroker "github.com/bryanArroyave/golang-utils/events/adapter/routerBroker"
	"github.com/bryanArroyave/golang-utils/events/enums"
	"github.com/bryanArroyave/golang-utils/events/factory"
	postgresdtos "github.com/bryanArroyave/golang-utils/gorm/dtos"
	utilsports "github.com/bryanArroyave/golang-utils/gorm/ports"
	loggerenums "github.com/bryanArroyave/golang-utils/logger/enums"
	"github.com/bryanArroyave/golang-utils/server"
	serverdtos "github.com/bryanArroyave/golang-utils/server/dtos"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc                 fx.Lifecycle
	App                *app.App
	Server             *server.APIRestServer
	Handlers           []infraports.IHttpHandler       `group:"handlers"`
	SubscriberHandlers []infraports.ISubscriberHandler `group:"subscriber_handlers"`
}

func main() {

	config.LoadEnvironmentVariables(os.Getenv("ENV"))

	serverConfig := &serverdtos.APIRestServerConfigDTO{
		GlobalPrefix: "",
		Port:         os.Getenv("PORT"),
	}

	fx.New(
		fx.Provide(
			context.Background,

			func(app *app.App) userports.IUserEventsRepository {
				return events.NewDomainEventAdapter(app.GetMessageBroker("domainBroker").Publisher)
			},
			func(app *app.App) categoryports.ICategoryRepository {
				return categoryadapters.NewCategoryAdapter(app.GetPostgresConnection("financial"))
			},
			func(app *app.App) categoryports.ICategoryService {
				return categoryadapters.NewCategoryAdapter(app.GetPostgresConnection("financial"))
			},

			func(app *app.App) userports.IUserService {
				return adapters.NewUserAdapter(app.GetPostgresConnection("financial"))
			},
			func(app *app.App) userports.IUserRepository {
				return adapters.NewUserAdapter(app.GetPostgresConnection("financial"))
			},
			func(app *app.App) *serverdtos.APIRestServerConfigDTO {
				serverConfig.App = app
				return serverConfig
			},
			server.NewAPIRestServer,
			func() *appdtos.LoggerConfigDTO {
				return &appdtos.LoggerConfigDTO{
					LoggerType:  loggerenums.Zerolog,
					ServiceName: os.Getenv("SERVICE_NAME"),
				}
			},
			func(config *appdtos.LoggerConfigDTO) *app.App {
				_app := app.NewApp(config)
				_app.AddPostgresConnection("financial", &postgresdtos.ConnectionDTO{
					URI:        os.Getenv("CONTROL_FINANCIAL_DB_URI"),
					Env:        os.Getenv("ENV"),
					MaxRetries: 3,
				})
				_app.AddMessageBroker("domainBroker", enums.Channels, &factory.FactoryConfig{
					GoChannel: &eventbroker.GoChannelConfig{
						BufferSize: 100,
					},
				})

				return _app
			},
			func(app *app.App) utilsports.IDBManager {
				return app.GetPostgresConnection("financial")
			},
		),
		subscribers.CategorySubscriberModule,
		auth.UsecasesModule,
		authhandlers.AuthModule,
		category.UsecasesModule,
		categoryhandlers.CategoryModule,
		fx.Invoke(
			setLifeCycle,
		),
	).Run()
}

func setLifeCycle(p Params) {
	p.Lc.Append(fx.Hook{
		OnStart: func(context.Context) error {

			err := migrations.StartMigrations(p.App.GetPostgresConnection("financial"))

			if err != nil {
				return err
			}

			p.Server.GetPrivateGroup().Use(middlewares.SessionCookie)

			for _, h := range p.Handlers {
				h.RegisterRoutes(p.Server.GetPublicGroup(), p.Server.GetPrivateGroup())
			}

			router := routerbroker.NewRouter()
			for _, h := range p.SubscriberHandlers {
				s := p.App.GetMessageBroker("domainBroker")
				h.RegisterRoutes(router, s.Subscriber)
			}

			go func() {
				if err := router.Run(context.Background()); err != nil {
					panic(err)
				}
			}()

			go func() {
				p.Server.RunServer()
			}()

			return nil
		},

		OnStop: func(ctx context.Context) error {

			return nil
		},
	})
}
