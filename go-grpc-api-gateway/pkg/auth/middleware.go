package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth/pb"
)

type AuthMiddlewareConfig struct {
	svc *ServiceClient
}

func InitAuthMiddleware(svc *ServiceClient) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{svc}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *fiber.Ctx) error {
	authorization := ctx.Get("Authorization")

	if authorization == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	token := strings.Split(authorization, "Bearer ")

	if len(token) < 2 {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	res, err := c.svc.Client.Validate(context.Background(), &pb.ValidateRequest{
		Token: token[1],
	})

	if err != nil || res.Status != http.StatusOK {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}

	userIdStr := strconv.Itoa(int(res.UserId))
	ctx.Set("userId", userIdStr)

	return ctx.Next()
}
