package auth

import (
	"fmt"

	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/auth/pb"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceCient struct {
	Client pb.AuthServiceClient
}

func InitServiceCient(c *config.Config) pb.AuthServiceClient {
	// using withInsecure because no SS is running
	cc,  err := grpc.Dial(c.AuthSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)

	}

	return pb.NewAuthServiceClient(cc)
}

