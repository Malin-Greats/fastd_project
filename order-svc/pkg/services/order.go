package services

import (
	"context"
	"net/http"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/client"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/db"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/models"
	"github.com/Malin-Greats/fastd_project/order-svc/pkg/pb"
)

type Server struct {
	H db.Handler
	ParcelSvc client.ParcelServiceClient
	pb.UnimplementedOrderServiceServer
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	parcel, err := s.ParcelSvc.FindOne(req.ParcelId)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil

	} else if parcel.Status >= http.StatusNotFound {
		return &pb.CreateOrderResponse{Status: parcel.Status, Error: parcel.Error}, nil
	} else if parcel.Data.DeliveryStage == "on_hold" {
		return &pb.CreateOrderResponse{Status: http.StatusConflict, Error: "Parcel not in move state"}, nil
	}

	order := models.Order{
		Price: parcel.Data.Cost,
		ParcelId: parcel.Data.Id,
		UserId: req.UserId,
	}
	s.H.DB.Create(&order)

	res, err := s.ParcelSvc.ChangeDeliveryStage(req.ParcelId, order.Id)

	if err != nil {
		return &pb.CreateOrderResponse{Status: http.StatusBadRequest, Error: err.Error()}, nil
	} else if res.Status == http.StatusConflict {
		s.H.DB.Delete(&models.Order{}, order.Id)

		return &pb.CreateOrderResponse{
			Status: http.StatusConflict, Error: res.Error}, nil
	}

	return &pb.CreateOrderResponse{
		Status: http.StatusCreated,
		Id: order.Id,
	}, nil

}