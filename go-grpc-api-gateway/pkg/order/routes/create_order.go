package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	order_request "github.com/raihantf/go-grpc-api-gateway/pkg/models/request"
	"github.com/raihantf/go-grpc-api-gateway/pkg/order/pb"
)

func CreateOrder(ctx *fiber.Ctx, c pb.OrderServiceClient) error {
	body := order_request.CreateOrderRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error parsing data")
	}

	userId := ctx.Locals("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ProductId: body.ProductId,
		Quantity:  body.Quantity,
		UserId:    userId.(int64),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "error bad gateway")
	}

	return ctx.JSON(&fiber.Map{
		"message": res,
	})
}
