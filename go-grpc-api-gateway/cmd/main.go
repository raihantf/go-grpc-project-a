package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth"
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"github.com/raihantf/go-grpc-api-gateway/pkg/order"
	"github.com/raihantf/go-grpc-api-gateway/pkg/product"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	app := fiber.New()

	authSvc := *auth.RegisterRoutes(app, &config)
	product.RegisterRoutes(app, &config, &authSvc)
	order.RegisterRoutes(app, &config, &authSvc)

	app.Listen(config.Port)
}
