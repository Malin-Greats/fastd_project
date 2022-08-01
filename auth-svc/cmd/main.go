package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/Malin-Greats/fastd_project/auth-svc/pkg/config"
	"github.com/Malin-Greats/fastd_project/auth-svc/pkg/db"
	

	"github.com/Malin-Greats/fastd_project/auth-svc/pkg/pb"
	"github.com/Malin-Greats/fastd_project/auth-svc/pkg/services"
	"github.com/Malin-Greats/fastd_project/auth-svc/pkg/utils"
	
)
func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	h := db.Init(c.DBUrl)

	jwt := utils.JwtWrapper{
		SecretKey: c.JWTSecretKey,
		Issuer: "auth-svc",
		ExpirationHours: 24 * 365,
	}


	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Auth Svc on", c.Port)
	
	s := services.Server{
		H: h,
		Jwt: jwt,
	}

	grpcServer := grpc.NewServer()


	pb.RegisterAuthServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}