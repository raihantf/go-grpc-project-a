package order

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth"
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"github.com/raihantf/go-grpc-api-gateway/pkg/order/routes"
)

func RegisterRoutes(app *fiber.App, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := app.Group("/order")
	routes.Use(a.AuthRequired)
	routes.Post("/", svc.CreateOrder)
}

func (svc *ServiceClient) CreateOrder(ctx *fiber.Ctx) error {
	err := routes.CreateOrder(ctx, svc.Client)
	if err != nil {
		panic(err)
	}
	return nil
}
