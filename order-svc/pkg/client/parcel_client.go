package client

import (
	"context"
	"fmt"

	"github.com/Malin-Greats/fastd_project/order-svc/pkg/pb"
	"google.golang.org/grpc"
)

type ParcelServiceClient struct {
	Client pb.ParcelServiceClient
}

func InitParcelServiceClient(url string) ParcelServiceClient {

	cc, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	c := ParcelServiceClient{
		Client: pb.NewParcelServiceClient(cc),
	}
	return c
}

func (c *ParcelServiceClient) FindOne(parcelId int64) (*pb.FindOneResponse, error) {
	req := &pb.FindOneRequest{
		Id: parcelId,
	}
	return c.Client.FindOne(context.Background(), req)
}


func (c *ParcelServiceClient) ChangeDeliveryStage(parcelId int64, orderId int64) (*pb.DeliveryStageResponse, error) {
	req := &pb.DeliveryStageRequest{
		Id: parcelId,
		OrderId: orderId,
	}

	return c.Client.DeliveryStage(context.Background(), req)
}