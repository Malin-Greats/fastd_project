package routes 

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/parcel/pb"
)


func FindOne(ctx *gin.Context, c pb.ParcelServiceClient) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 32)


	res, err := c.FindOne(context.Background(), &pb.FindOneRequest{
		Id: int64(id),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}