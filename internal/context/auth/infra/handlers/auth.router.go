package handlers

import (
	"os"

	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	applicationports "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/ports"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type handler struct {
	registerUserUsecase   applicationports.IRegisterUser
	getUserByEmailUsecase applicationports.IGetUserByEmail

	googleOauth2Config oauth2.Config
	githuboauth2Config oauth2.Config
	oauth2StateString  string
}

type Result struct {
	fx.Out

	Handler infraports.IHttpHandler `group:"handlers"`
}

func New(
	registerUserUsecase applicationports.IRegisterUser,
	getUserByEmailUsecase applicationports.IGetUserByEmail,
) Result {
	return Result{
		Handler: &handler{
			registerUserUsecase:   registerUserUsecase,
			getUserByEmailUsecase: getUserByEmailUsecase,

			googleOauth2Config: oauth2.Config{
				ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
				ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("GOOGLE_URL_REDIRECT"),
				Scopes:       []string{"openid", "email", "profile"},
				Endpoint:     google.Endpoint,
			},
			oauth2StateString: "random",
			githuboauth2Config: oauth2.Config{
				ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
				ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("GITHUB_URL_REDIRECT"),
				Scopes:       []string{"read:user", "user:email"},
				Endpoint:     github.Endpoint,
			},
		},
	}
}

type GithubUser struct {
	ID        int64  `json:"id"`
	Login     string `json:"login"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
}

type GithubUserEmails []*GithubUserEmail

func (g *GithubUserEmails) GetPrimaryEmail() string {
	for _, email := range *g {
		if email.Primary {
			return email.Email
		}
	}
	return ""
}

type GithubUserEmail struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

type GoogleUser struct {
	ID            string `json:"sub"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"email_verified"`
	Picture       string `json:"picture"`
	Name          string `json:"name"`
}

type LoggedUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Img   string `json:"img"`
	// TODO: meter en un enum
	Provider string `json:"provider"`
}

func (h *handler) RegisterRoutes(publicGroup *echo.Group, privateGroup *echo.Group) {

	authGroup := publicGroup.Group("/auth")

	authGroup.GET("/login", h.Login)
	authGroup.GET("/callback", h.Callback)

}

var AuthModule = fx.Module("category", fx.Provide(
	New,
))
