package transport

import (
	"auth_microservice/internal/domain"
	"context"
	"github.com/gin-gonic/gin"
	"time"
)

type Users interface {
	SignIn(ctx context.Context, inp domain.SignInInput) (string, string, error)
	ParseToken(token string) (int, error)
	RefreshTokens(ctx context.Context, refreshToken string) (string, string, error)
	GetRefreshTokenTTL() time.Duration
}

type Handler struct {
	usersService Users
}

func NewHandler(users Users) *Handler {
	return &Handler{
		usersService: users,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	usersApi := router.Group("/auth")
	{
		usersApi.POST("/sign-in", h.signIn)
		usersApi.GET("/refresh", h.refresh)
	}

	return router
}
