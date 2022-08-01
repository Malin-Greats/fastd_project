package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/config"
	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/db"
	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/pb"
	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/services"
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

	fmt.Println("Parcel Svc on", c.Port)
	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterParcelServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}


}