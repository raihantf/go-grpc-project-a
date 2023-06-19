package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(app *fiber.App, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := app.Group("/auth")
	routes.Post("/register")
}
