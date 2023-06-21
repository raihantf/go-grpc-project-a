package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	product_request "github.com/raihantf/go-grpc-api-gateway/pkg/models/request"
	"github.com/raihantf/go-grpc-api-gateway/pkg/product/pb"
)

func CreateProduct(ctx *fiber.Ctx, c pb.ProductServiceClient) error {
	body := product_request.CreateProductRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error parsing data")
	}

	res, err := c.CreateProduct(context.Background(), &pb.CreateProductRequest{
		Name:  body.Name,
		Stock: body.Stock,
		Price: body.Price,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "error bad gateway")
	}

	return ctx.JSON(&fiber.Map{
		"message": res,
	})
}
