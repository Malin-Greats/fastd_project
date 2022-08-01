package db

import (
	"log"

	"github.com/Malin-Greats/fastd_project/parcel-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB

}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&models.Parcel{})
	db.AutoMigrate(&models.DeliveryStageLog{})

	return Handler{}
}