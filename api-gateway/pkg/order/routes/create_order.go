package routes

import (
	"context"
	"net/http"

	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/order/pb"
	"github.com/gin-gonic/gin"
)

type CreateOrderRequestBody struct {
	ParcelId int64 `json:"parcelId"`
	Size string `json:"size"`
	UserId int64 `json:"userId"`
}


func CreateOrder(ctx *gin.Context, c pb.OrderServiceClient) {
	body := CreateOrderRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	userId, _ := ctx.Get("userId")

	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		ParcelId: body.ParcelId,
		Size: body.Size,
		UserId: userId.(int64),
	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
		}

	ctx.JSON(http.StatusCreated, &res)
}