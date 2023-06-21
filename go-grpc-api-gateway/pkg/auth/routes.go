package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth/routes"
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
)

func RegisterRoutes(app *fiber.App, c *config.Config) *ServiceClient {
	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := app.Group("/auth")
	routes.Post("/register", svc.Register)
	routes.Post("/login", svc.Login)

	return svc
}

func (svc *ServiceClient) Register(ctx *fiber.Ctx) error {
	err := routes.Register(ctx, svc.Client)
	if err != nil {
		panic(err)
	}
	return nil
}

func (svc *ServiceClient) Login(ctx *fiber.Ctx) error {
	err := routes.Login(ctx, svc.Client)
	if err != nil {
		panic(err)
	}
	return nil
}
