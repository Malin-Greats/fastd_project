package routes
import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Malin-Greats/fastd_project/api-gateway/pkg/parcel/pb"

)

type CreateParcelRequestBody struct {
	Description		string 	`json:"description"`
    Sku				string 	`json:"sku"`
    Queque			int64 	`json:"queque"`
    Cost			int64 	`json:"cost"`
    Paddress		string 	`json:"paddress"`
    Daddress		string 	`json:"daddress"`
    Ptime			string 	`json:"ptime"`
    Owner			string 	`json:"owner"`
    Driver			string 	`json:"driver"`
}

func CreateParcel(ctx *gin.Context, c pb.ParcelServiceClient) {
	body := CreateParcelRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	res, err := c.CreateParcel(context.Background(), &pb.CreateParcelRequest{
		Description:	body.Description,
		Sku:			body.Sku,
		Queque:		body.Queque,
		Cost:		body.Cost,
		Paddress:	body.Paddress,
		Daddress:	body.Daddress,
		Ptime:		body.Ptime,
		Owner:		body.Owner,
		Driver:		body.Driver,

	})

	if err != nil {
		ctx.AbortWithError(http.StatusBadGateway, err)
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}