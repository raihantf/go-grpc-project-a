package routes

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth/pb"
	auth_request "github.com/raihantf/go-grpc-api-gateway/pkg/models/request"
)

func Login(ctx *fiber.Ctx, c pb.AuthServiceClient) error {
	body := auth_request.LoginRequest{}

	if err := ctx.BodyParser(body); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "error parsing data")
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	})

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, "error bad gateway")
	}

	return ctx.JSON(&fiber.Map{
		"message": res,
	})
}
