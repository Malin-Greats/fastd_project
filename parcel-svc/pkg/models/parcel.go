package models

type Parcel struct {
	Id					int64 				`json:"id" gorm:"primaryKey"`
	Name 				string 				`json:"name"`
	Description			string				`json:"description"` 
    Sku					string				`json:"sku"` 
    Queque				int64				`json:"queque"` 
    Cost				int64				`json:"cost"` 
    Pickup_address		string				`json:"paddress"` 
    Delivery_address	string				`json:"daddress"` 
    Pickup_time			string				`json:"ptime"` 
    Owner				string				`json:"owner"` 
    Driver				string				`json:"driver"`
	DeliveryStage 		DeliveryStageLog 	`gorm:"foreignKey:ParcelRefer"`
}
