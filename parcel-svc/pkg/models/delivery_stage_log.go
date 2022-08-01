package models

type DeliveryStageLog struct {
	Id				int64 `json:"id" gorm:"primaryKey"`
	OrderId 		int64 `json:"order_id"`
	ParcelRefer		int64 `json:"parcel_id"`
	Stage   		string `json:"stage"`
}


