package auth

import (
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/auth/routes"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, c *config.Config) *ServiceCient {
	svc := &ServiceCient {
		Client: InitServiceCient(c),
	}

	routes := r.Group("/auth")
	routes.POST("/register", svc.Register)
	routes.POST("/ogin", svc.Login)

	return svc
}

func (svc *ServiceCient) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.Client)
}

func (svc *ServiceCient) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.Client)
}