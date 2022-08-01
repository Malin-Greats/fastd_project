package order

import (
	"fmt"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/order/pb"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceCient struct {
	Client pb.OrderServiceClient
}

func InitServiceCient(c *config.Config) pb.OrderServiceClient {
	cc, err := grpc.Dial(c.OrderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect", err)
	}

	return pb.NewOrderServiceClient(cc)
}
