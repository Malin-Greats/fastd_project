package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/auth"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/order"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/parcel"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	parcel.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}