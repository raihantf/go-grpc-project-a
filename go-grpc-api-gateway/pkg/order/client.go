package order

import (
	config "github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"github.com/raihantf/go-grpc-api-gateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.OrderServiceClient
}

func InitServiceClient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	return pb.NewOrderServiceClient(cc)
}
