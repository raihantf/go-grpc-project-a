package auth

import (
	"github.com/raihantf/go-grpc-api-gateway/pkg/auth/pb"
	config "github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(c *config.Config) pb.AuthServiceClient {
	cc, err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return pb.NewAuthServiceClient(cc)
}
