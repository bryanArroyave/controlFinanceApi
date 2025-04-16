package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/application/usecases/registerUser/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/enums"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/auth/domain/services"
	infradtos "github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/auth/dtos"
	"github.com/bryanArroyave/eventsplit/back/user-service/internal/context/shared/infra/handlers/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
)

func (h *handler) Callback(c echo.Context) error {

	if c.QueryParam("state") != h.oauth2StateString {
		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidState", fmt.Errorf("missing state"), nil))
	}

	provider := c.QueryParam("provider")
	code := c.QueryParam("code")

	if provider == "" {
		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("missingProvider", fmt.Errorf("missing provider"), nil))
	}

	if code == "" {
		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("missingCode", fmt.Errorf("missing code"), nil))
	}

	var callbackFunc func(echo.Context, *oauth2.Token) error
	var config oauth2.Config

	if provider == enums.Github.String() {
		callbackFunc = h.callbackGithub
		config = h.githuboauth2Config
	} else if provider == enums.Google.String() {
		callbackFunc = h.callbackGoogle
		config = h.googleOauth2Config
	} else {
		return c.JSON(http.StatusUnauthorized, utils.BuildErrorResponse("invalidProvider", fmt.Errorf("invalid provider"), nil))
	}

	token, err := config.Exchange(c.Request().Context(), code)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("unexpectedError", err, nil))
	}

	return callbackFunc(c, token)
}

func (h *handler) callbackGoogle(c echo.Context, oauthToken *oauth2.Token) error {

	client := h.googleOauth2Config.Client(c.Request().Context(), oauthToken)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil || resp.StatusCode != http.StatusOK {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to get user info"})
	}
	defer resp.Body.Close()

	var user *GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "invalid user response"})
	}

	LoggedUser := &LoggedUser{
		Email:    user.Email,
		Name:     user.Name,
		Provider: enums.Google.String(),
		Img:      user.Picture,
	}

	cookie, err := h.getCookie(c.Request().Context(), LoggedUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("error", err, nil))
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, utils.BuildSuccessResponse(echo.Map{"access_token": cookie.Value}, nil))
}

func (h *handler) callbackGithub(c echo.Context, oauthToken *oauth2.Token) error {

	client := h.githuboauth2Config.Client(c.Request().Context(), oauthToken)
	data, err := client.Get("https://api.github.com/user")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("failedToGetUserData", err, nil))
	}
	defer data.Body.Close()

	emails, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.BuildErrorResponse("failedToGetUserData", err, nil))
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

	LoggedUser := &LoggedUser{
		Email:    githubUserEmail.GetPrimaryEmail(),
		Name:     githubUser.Login,
		Provider: enums.Github.String(),
		Img:      githubUser.AvatarURL,
	}

	cookie, err := h.getCookie(c.Request().Context(), LoggedUser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "failed to generate JWT"})
	}

	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, utils.BuildSuccessResponse(echo.Map{"access_token": cookie.Value}, nil))

}

func (h *handler) getCookie(ctx context.Context, loggedUser *LoggedUser) (*http.Cookie, error) {

	userID, err := h.verifyUser(ctx, loggedUser)
	if err != nil {
		return nil, err
	}

	claims := &infradtos.Claims{
		Email:  loggedUser.Email,
		UserID: userID,
	}

	tokenService := services.NewTokenService()
	accessToken, expiredAt, err := tokenService.GenerateToken(claims)
	if err != nil {
		return nil, err
	}

	cookie := new(http.Cookie)
	cookie.Name = "control_finance_session_token"
	cookie.Value = accessToken
	cookie.HttpOnly = true
	// cookie.Secure = true // Solo si usas HTTPS
	cookie.SameSite = http.SameSiteLaxMode
	cookie.Path = "/"
	cookie.Expires = expiredAt // Igual que el tiempo del token

	return cookie, nil
}

func (h *handler) verifyUser(ctx context.Context, loggedUser *LoggedUser) (int, error) {

	user, err := h.getUserByEmailUsecase.GetUserByEmail(ctx, loggedUser.Email)

	if err != nil {
		return 0, err
	}

	var userID int
	if user == nil {

		userToRegister := &dtos.RegisterUserParam{
			Name:     loggedUser.Name,
			Email:    loggedUser.Email,
			Provider: loggedUser.Provider,
			Img:      loggedUser.Img,
		}
		userID, err = h.registerUserUsecase.RegisterUser(ctx, userToRegister)
		if err != nil {
			return 0, err
		}

	} else {
		userID = int(user.ID)
	}
	return userID, nil
}
