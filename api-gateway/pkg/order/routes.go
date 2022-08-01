package order

import (
	"github.com/gin-gonic/gin"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/auth"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/order/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceCient) {

	a := auth.InitAuthMiddleware(authSvc)

	svc := &ServiceCient {
		Client: InitServiceCient(c),
	}

	routes := r.Group("/order")
	routes.Use(a.AuthRequired)
	routes.POST("/", svc.CreateOrder)
}

func (svc *ServiceCient) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.Client)
}
 