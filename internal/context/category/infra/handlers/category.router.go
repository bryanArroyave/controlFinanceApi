package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	infraports "github.com/bryanArroyave/eventsplit/back/user-service/infra/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/category/application/ports"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/auth/middlewares"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

type handler struct {
	addCategoryUsecase    ports.IAddCategory
	addSubcategoryUsecase ports.ISaveSubcategory
}

type Result struct {
	fx.Out

	Handler infraports.IHttpHandler `group:"handlers"`
}

func New(
	addCategoryUsecase ports.IAddCategory,
	addSubcategoryUsecase ports.ISaveSubcategory,
) Result {
	return Result{
		Handler: &handler{
			addCategoryUsecase:    addCategoryUsecase,
			addSubcategoryUsecase: addSubcategoryUsecase,
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
type OAuthToken struct {
	ID          uint  `gorm:"primaryKey"`
	GithubID    int64 `gorm:"uniqueIndex"`
	Username    string
	AccessToken string
	TokenType   string
	Email       string
	Expiry      time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (h *handler) RegisterRoutes(e *echo.Echo) {

	oauth2Config := oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),     // Reemplaza con tu Client ID
		ClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"), // Reemplaza con tu Client Secret
		RedirectURL:  os.Getenv("GITHUB_URL_REDIRECT"),  // URL de redirección
		Scopes:       []string{"read:user", "user:email"},
		Endpoint:     github.Endpoint, // El endpoint de OAuth para GitHub
	}
	oauth2StateString := "random"

	googleOAuthConfig := oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),     // Reemplaza con tu Client ID
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"), // Reemplaza con tu Client Secret
		RedirectURL:  os.Getenv("GOOGLE_URL_REDIRECT"),
		Scopes:       []string{"openid", "email", "profile"},
		Endpoint:     google.Endpoint,
	}

	type GoogleUser struct {
		ID            string `json:"sub"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"email_verified"`
		Picture       string `json:"picture"`
		Name          string `json:"name"`
	}

	userGroup := e.Group("/category")

	userGroup.POST("", h.AddCategory)
	userGroup.POST("/:category_id/subcategory", h.AddSubcategory, middlewares.JWTMiddleware)
	userGroup.GET("/:category_id/subcategory", func(c echo.Context) error {
		categoryID := c.Param("category_id")
		return c.String(http.StatusOK, "Category ID: "+categoryID)
	}, middlewares.JWTMiddleware)

	authGroup := e.Group("/auth")

	authGroup.GET("/login/github", func(c echo.Context) error {
		url := oauth2Config.AuthCodeURL(oauth2StateString, oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusFound, url)
	})

	authGroup.GET("/callback/github", func(c echo.Context) error {

		if c.QueryParam("state") != oauth2StateString {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid state",
			})
		}

		code := c.QueryParam("code")
		if code == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "code not found",
			})
		}

		// Intercambia el código por un token
		token, err := oauth2Config.Exchange(c.Request().Context(), code)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to exchange token")
		}

		// Usa el token para hacer peticiones a la API de GitHub
		client := oauth2Config.Client(c.Request().Context(), token)
		data, err := client.Get("https://api.github.com/user")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to get user data")
		}
		defer data.Body.Close()

		emails, err := client.Get("https://api.github.com/user/emails")
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to get user data")
		}
		defer emails.Body.Close()

		var githubUser *GithubUser
		if err := json.NewDecoder(data.Body).Decode(&githubUser); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid user response"})
		}
		var githubUserEmail GithubUserEmails
		if err := json.NewDecoder(emails.Body).Decode(&githubUserEmail); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid user response"})
		}

		tokenEntry := OAuthToken{
			GithubID:    githubUser.ID,
			Username:    githubUser.Login,
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
			Expiry:      token.Expiry,
			Email:       githubUserEmail.GetPrimaryEmail(),
		}

		fmt.Println(tokenEntry)

		t, err := middlewares.GenerateJWT(tokenEntry.Username, tokenEntry.Email)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate JWT"})
		}

		cookie := new(http.Cookie)
		cookie.Name = "control_finance_session_token"
		cookie.Value = t
		cookie.HttpOnly = true
		// cookie.Secure = true // Solo si usas HTTPS
		cookie.SameSite = http.SameSiteLaxMode
		cookie.Path = "/"
		cookie.Expires = time.Now().Add(24 * time.Hour) // Igual que el tiempo del token

		c.SetCookie(cookie)

		// Procesa la respuesta de GitHub
		if data.StatusCode == http.StatusOK {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"access_token":  token.AccessToken,
				"t":             t,
				"refresh_token": token.RefreshToken,
				"token_type":    token.TokenType,
				"expiry":        token.Expiry,
			})
		} else {
			return c.String(http.StatusUnauthorized, "Failed to authenticate with GitHub")
		}

		//

		// Obtén el código de autorización de la URL

		// Intercambia el código por un token

		// Devuelve el token generado

	})

	authGroup.GET("/login/google", func(c echo.Context) error {
		url := googleOAuthConfig.AuthCodeURL(oauth2StateString, oauth2.AccessTypeOffline)
		return c.Redirect(http.StatusFound, url)
	})
	authGroup.GET("/callback/google", func(c echo.Context) error {
		if c.QueryParam("state") != oauth2StateString {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid state"})
		}

		code := c.QueryParam("code")
		if code == "" {
			return c.JSON(http.StatusUnauthorized, echo.Map{"error": "code not found"})
		}

		token, err := googleOAuthConfig.Exchange(c.Request().Context(), code)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to exchange token"})
		}

		client := googleOAuthConfig.Client(c.Request().Context(), token)
		resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil || resp.StatusCode != http.StatusOK {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get user info"})
		}
		defer resp.Body.Close()

		var user *GoogleUser
		if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid user response"})
		}

		// Guardar en la DB (misma tabla o una nueva según tu modelo)
		tokenEntry := OAuthToken{
			GithubID:    0, // puedes usar otro campo como `ProviderID` o `Email`
			Username:    user.Email,
			AccessToken: token.AccessToken,
			TokenType:   token.TokenType,
			Expiry:      token.Expiry,
		}
		fmt.Println(tokenEntry)

		return c.JSON(http.StatusOK, echo.Map{
			"access_token": token.AccessToken,
			"user":         user.Email,
		})
	})

}

var CategoryModule = fx.Module("category", fx.Provide(
	New,
))
