package product

import (
	"github.com/raihantf/go-grpc-api-gateway/pkg/config"
	"github.com/raihantf/go-grpc-api-gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ProductServiceClient
}

func InitServiceClient(c *config.Config) pb.ProductServiceClient {
	cc, err := grpc.Dial(c.ProductSvcUrl, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return pb.NewProductServiceClient(cc)
}
