package rest

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/KimNattanan/exprec-backend/internal/user/dto"
	"github.com/KimNattanan/exprec-backend/internal/user/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type HttpUserHandler struct {
	userUseCase       usecase.UserUseCase
	googleOauthConfig *oauth2.Config
}

func NewHttpUserHandler(useCase usecase.UserUseCase, clientID, clientSecret, redirectURL string) *HttpUserHandler {
	return &HttpUserHandler{
		userUseCase: useCase,
		googleOauthConfig: &oauth2.Config{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			RedirectURL:  redirectURL,
			Scopes:       []string{"openid", "email", "profile"},
			Endpoint:     google.Endpoint,
		},
	}
}

func (h *HttpUserHandler) GetUser(c *fiber.Ctx) error {
	user_id, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}

	user, err := h.userUseCase.FindByID(user_id)
	if err != nil {
		return responses.Error(c, err)
	}

	return c.JSON(dto.ToUserResponse(user))
}

func (h *HttpUserHandler) GoogleLogin(c *fiber.Ctx) error {
	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  time.Now().Add(10 * time.Minute),
		HTTPOnly: true,
	})
	url := h.googleOauthConfig.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce, oauth2.SetAuthURLParam("prompt", "consent select_account"))
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

func (h *HttpUserHandler) GoogleCallback(c *fiber.Ctx) error {
	state := c.Cookies("oauthstate")
	if c.Query("state") != state {
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "invalid oauth state")
	}

	code := c.Query("code")
	if code == "" {
		return responses.ErrorWithMessage(c, appError.ErrInvalidData, "code not found")
	}

	token, err := h.googleOauthConfig.Exchange(c.Context(), code)
	if err != nil {
		log.Printf("failed to exchange token: %v\n", err)
		return responses.ErrorWithMessage(c, appError.ErrUnauthorized, "failed to exchange token")
	}

	client := h.googleOauthConfig.Client(c.Context(), token)
	res, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return responses.ErrorWithMessage(c, err, "failed to get user info")
	}
	defer res.Body.Close()

	var userInfo map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&userInfo); err != nil {
		return responses.ErrorWithMessage(c, err, "failed to decode user info")
	}

	jwtToken, _, err := h.userUseCase.LoginOrRegisterWithGoogle(userInfo, token)
	if err != nil {
		return responses.Error(c, err)
	}

	isProd := os.Getenv("ENV") == "production"

	c.Cookie(&fiber.Cookie{
		Name:     "oauthstate",
		Expires:  time.Now(),
		HTTPOnly: true,
		Secure:   false,
	})
	c.Cookie(&fiber.Cookie{
		Name:     "loginToken",
		Value:    jwtToken,
		Expires:  time.Now().Add(14 * 24 * time.Hour),
		HTTPOnly: true,
		Secure:   isProd,
		SameSite: "None",
		Path:     "/",
		Domain:   ".exprec.kim",
	})

	return c.Redirect(os.Getenv("FRONTEND_OAUTH_REDIRECT_URL"), fiber.StatusSeeOther)
}

func (h *HttpUserHandler) Logout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "loginToken",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
	})
	return responses.Message(c, fiber.StatusOK, "logged out successfully")
}

func (h *HttpUserHandler) Delete(c *fiber.Ctx) error {
	user_id, err := uuid.Parse(c.Locals("user_id").(string))
	if err != nil {
		return responses.Error(c, appError.ErrInvalidData)
	}
	if err := h.userUseCase.Delete(user_id); err != nil {
		return responses.Error(c, err)
	}
	return responses.Message(c, fiber.StatusOK, "user deleted")
}
