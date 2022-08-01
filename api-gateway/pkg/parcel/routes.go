package parcel

import (
	"github.com/gin-gonic/gin"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/auth"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/parcel/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceCient) {
	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceCient{
		Client: InitServiceCient(c),
	}

	routes := r.Group("/parcel")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateParcel)
	routes.GET("/", svc.FindOne)
}

func (svc *ServiceCient) FindOne(ctx *gin.Context) {
	routes.FindOne(ctx, svc.Client)
}

func (svc *ServiceCient) CreateParcel(ctx *gin.Context) {
	routes.CreateParcel(ctx, svc.Client)
}