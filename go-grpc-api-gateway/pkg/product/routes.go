package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth"
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"github.com/raihantf/go-grpc-api-gateway/pkg/product/routes"
)

func RegisterRoutes(app *fiber.App, c *config.Config, authSvc *auth.ServiceClient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceClient{
		Client: InitServiceClient(c),
	}

	routes := app.Group("/product")
	routes.Use(a.AuthRequired)
	routes.Post("/", svc.CreateProduct)
	routes.Get("/:id", svc.FindOne)
}

func (svc *ServiceClient) FindOne(ctx *fiber.Ctx) error {
	err := routes.FindOne(ctx, svc.Client)
	if err != nil {
		panic(err)
	}
	return nil
}

func (svc *ServiceClient) CreateProduct(ctx *fiber.Ctx) error {
	err := routes.CreateProduct(ctx, svc.Client)
	if err != nil {
		panic(err)
	}
	return nil
}
