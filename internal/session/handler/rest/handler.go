package rest

import (
	"os"
	"time"

	"github.com/KimNattanan/exprec-backend/internal/session/dto"
	"github.com/KimNattanan/exprec-backend/internal/session/usecase"
	appError "github.com/KimNattanan/exprec-backend/pkg/apperror"
	"github.com/KimNattanan/exprec-backend/pkg/responses"
	"github.com/KimNattanan/exprec-backend/pkg/token"
	"github.com/gofiber/fiber/v2"
)

type HttpSessionHandler struct {
	tokenMaker     *token.JWTMaker
	sessionUseCase usecase.SessionUseCase
}

func NewHttpSessionHandler(useCase usecase.SessionUseCase, secretKey string) *HttpSessionHandler {
	return &HttpSessionHandler{
		sessionUseCase: useCase,
		tokenMaker:     token.NewJWTMaker(secretKey),
	}
}

func removeToken(c *fiber.Ctx) {
	domain := ""
	isProd := os.Getenv("ENV") == "production"
	if isProd {
		domain = ".exprec.kim"
	}
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "",
		Expires:  time.Now(),
		HTTPOnly: true,
		Domain:   domain,
	})
}

func (h *HttpSessionHandler) RenewToken(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	claims, err := h.tokenMaker.VerfiyToken(tokenStr)
	if err != nil {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}
	session, err := h.sessionUseCase.FindByID(claims.RegisteredClaims.ID)
	if err != nil {
		removeToken(c)
		return responses.Error(c, appError.ErrInternalServer)
	}
	if session.IsRevoked {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}
	if session.UserEmail != claims.Email {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}

	accessToken, accessClaims, err := h.tokenMaker.CreateToken(claims.ID, claims.Email, 15*time.Minute)
	if err != nil {
		return responses.Error(c, appError.ErrInternalServer)
	}

	return c.JSON(&dto.RenewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessClaims.RegisteredClaims.ExpiresAt.Time,
	})
}

func (h *HttpSessionHandler) Logout(c *fiber.Ctx) error {
	tokenStr := c.Cookies("token")
	claims, err := h.tokenMaker.VerfiyToken(tokenStr)
	if err != nil {
		removeToken(c)
		return responses.Error(c, appError.ErrUnauthorized)
	}
	if err := h.sessionUseCase.Delete(claims.RegisteredClaims.ID); err != nil {
		return responses.Error(c, appError.ErrInternalServer)
	}
	removeToken(c)
	return responses.Message(c, fiber.StatusOK, "logged out successfully")
}
