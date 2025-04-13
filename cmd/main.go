package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bryanArroyave/eventsplit/back/user-service/infra/migrations"
	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/infra/seeds"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/domain/ports"
	categoryadapters "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/adapters/postgres"
	categoryhandlers "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/infra/handlers"
	"github.com/bryanArroyave/golang-utils/app"
	appdtos "github.com/bryanArroyave/golang-utils/app/dtos"
	postgresdtos "github.com/bryanArroyave/golang-utils/gorm/dtos"
	utilsports "github.com/bryanArroyave/golang-utils/gorm/ports"
	"github.com/bryanArroyave/golang-utils/logger/enums"
	"github.com/bryanArroyave/golang-utils/server"
	serverdtos "github.com/bryanArroyave/golang-utils/server/dtos"
	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

type Params struct {
	fx.In

	Lc       fx.Lifecycle
	App      *app.App
	Server   *server.APIRestServer
	Handlers []infraports.IHttpHandler `group:"handlers"`
}

func main() {

	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	fmt.Println("1", os.Getenv("GITHUB_CLIENT_ID"))
	fmt.Println("1", os.Getenv("GITHUB_CLIENT_SECRET"))

	serverConfig := &serverdtos.APIRestServerConfigDTO{
		GlobalPrefix: "",
		Port:         os.Getenv("PORT"),
	}

	fx.New(

		fx.Provide(
			context.Background,
			func(app *app.App) ports.ICategoryRepository {
				return categoryadapters.NewCategoryAdapter(app.GetPostgresConnection("financial"))
			},
			func(app *app.App) ports.ICategoryService {
				return categoryadapters.NewCategoryAdapter(app.GetPostgresConnection("financial"))
			},
			func(app *app.App) *serverdtos.APIRestServerConfigDTO {
				serverConfig.App = app
				return serverConfig
			},
			server.NewAPIRestServer,
			func() *appdtos.LoggerConfigDTO {
				return &appdtos.LoggerConfigDTO{
					LoggerType:  enums.Zerolog,
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

				return _app
			},
			func(app *app.App) utilsports.IDBManager {
				return app.GetPostgresConnection("financial")
			},
		),
		// user.UsecasesModule,
		// userhandlers.UserModule,
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

			seed := seeds.NewControlFinancialSeed(p.App.GetPostgresConnection("financial"))
			seed.Exec()

			for _, h := range p.Handlers {
				h.RegisterRoutes(p.Server.GetEchoInstance())
			}

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
