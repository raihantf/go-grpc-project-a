package routes

import (
	"context"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/product/pb"
)

func FindOne(ctx *fiber.Ctx, c pb.ProductServiceClient) error {
	id, _ := strconv.ParseInt(ctx.Params("id"), 10, 32)

	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "error bad gateway")
	}

	return ctx.JSON(&fiber.Map{
		"message": res,
	})
}
