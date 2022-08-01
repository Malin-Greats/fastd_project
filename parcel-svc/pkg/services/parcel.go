package services
import (
	"context"
	"net/http"

	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/db"
	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/models"
	pb "github.com/Malin-Greats/fastd_project/parcel-svc/pkg/pb"
	
)

type Server struct {
	H db.Handler
	pb.UnimplementedParcelServiceServer
}

func (s *Server) CreateParcel(ctx context.Context, req *pb.CreateParcelRequest) (*pb.CreateParcelResponse, error) {
	var parcel models.Parcel
	parcel.Description		=	req.Description
	parcel.Sku				=	req.Sku
	parcel.Queque			=	req.Queque
	parcel.Cost				=	req.Cost
	parcel.Pickup_address	=	req.PickupAddress
	parcel.Delivery_address	=	req.DeliveryAddress
	parcel.Pickup_time		=	req.PickupTime
	parcel.Owner			=	req.Owner
	parcel.Driver			=	req.Driver

	if result := s.H.DB.Create(&parcel); result.Error != nil {
		return &pb.CreateParcelResponse{
			Status: http.StatusConflict,
			Error: result.Error.Error(),
		}, nil
	}

	return &pb.CreateParcelResponse{
		Status: http.StatusCreated,
		Id: parcel.Id,
	}, nil
}

func (s *Server) FindOne(ctx context.Context, req *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	var parcel models.Parcel

	if result := s.H.DB.First(&parcel, req.Id); result.Error != nil {
		return &pb.FindOneResponse{
			Status: http.StatusNotFound,
			Error: result.Error.Error(),		
		}, nil
	}

	data := &pb.FindOneData{
		Id: 				parcel.Id,
		Description: 		parcel.Description,
		Sku:				parcel.Sku,		
		Queque:				parcel.Queque,	
		Cost:				parcel.Cost,		
		PickupAddress:		parcel.Pickup_address,
		DeliveryAddress:	parcel.Delivery_address,
		PickupTime:			parcel.Pickup_time,
		Owner:				parcel.Owner,	
		Driver:				parcel.Driver,
	}

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data: data,	}, nil
}


func (s *Server) ChangeDeliveryStage(ctx context.Context, req *pb.DeliveryStageRequest) (*pb.DeliveryStageResponse, error) {
	var parcel models.Parcel

	if result := s.H.DB.First(&parcel, req.Id); result.Error != nil {
		return &pb.DeliveryStageResponse{
			Status: http.StatusNotFound,
			Error: result.Error.Error(),
		}, nil
	}

	if parcel.DeliveryStage.Stage == "on_hold" {
		return &pb.DeliveryStageResponse{
			Status: http.StatusConflict,
			Error: "Parcel not in moving state",
		}, nil
	}
	var log models.DeliveryStageLog
	if result := s.H.DB.Where(&models.DeliveryStageLog{OrderId: req.OrderId}).First(&log); result.Error == nil {
		return &pb.DeliveryStageResponse{
			Status: http.StatusConflict,
			Error: "Parcel already in the request stage",
		}, nil
	}

	parcel.DeliveryStage.Stage = req.DeliveryStage
	s.H.DB.Save(&parcel)

	log.OrderId = req.OrderId
	log.ParcelRefer = parcel.Id

	s.H.DB.Create(&log)

	return &pb.DeliveryStageResponse{
		Status: http.StatusOK,
	}, nil
	
}
