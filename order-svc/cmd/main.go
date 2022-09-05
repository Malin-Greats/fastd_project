package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Malin-Greats/fastd_project/order-svc/pkg/client"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/config"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/db"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/pb"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/services"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	parcelSvc := client.InitParcelServiceClient(c.ParcelSvcUrl)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Order Svc on", c.Port)

	s := services.Server{
		H: h,
		ParcelSvc: parcelSvc,
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}