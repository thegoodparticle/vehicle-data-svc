package store

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/thegoodparticle/vehicle-data-svc/models"
)

var owners = []models.Owner{
	{
		OwnerID:      "KA20123456",
		FirstName:    "John",
		LastName:     "Doe",
		MobileNumber: "8722312345",
		AadharNumber: "123412341234",
		Address:      "Banashankari, 3rd Stage, Benagluru - 560012",
	},
	{
		OwnerID:      "KA20231211",
		FirstName:    "Tim",
		LastName:     "David",
		MobileNumber: "9422312345",
		AadharNumber: "121234345656",
		Address:      "1st Cross, Bellandur, Benagluru - 560009",
	},
}

var vehicles = []models.Vehicle{
	{
		RegistrationNumber:     "KA20AB1234",
		VehicleModel:           "Bolero",
		Company:                "Mahindra",
		EngineNumber:           "52WVC10338",
		ChassisNumber:          "QWERTY12321233ABC",
		OwnerID:                "KA20123456",
		RegistrationIssueDate:  time.Date(2010, 5, 17, 20, 34, 58, 651387237, time.UTC),
		RegistrationExpiryDate: time.Date(2042, 5, 17, 20, 34, 58, 651387237, time.UTC),
	},
	{
		RegistrationNumber:     "KA20CD5678",
		VehicleModel:           "Polo",
		Company:                "VW",
		EngineNumber:           "62WVC10388",
		ChassisNumber:          "ASDFGH12321233CDE",
		OwnerID:                "KA20231211",
		RegistrationIssueDate:  time.Date(2015, 1, 12, 15, 20, 52, 561387237, time.UTC),
		RegistrationExpiryDate: time.Date(2047, 1, 12, 15, 20, 52, 561387237, time.UTC),
	},
	{
		RegistrationNumber:     "KA20EF9012",
		VehicleModel:           "Suzuki",
		Company:                "Maruti",
		EngineNumber:           "62WVC10398",
		ChassisNumber:          "ASDFGH12321233EFG",
		OwnerID:                "KA20231211",
		RegistrationIssueDate:  time.Date(2015, 1, 10, 15, 20, 52, 561387237, time.UTC),
		RegistrationExpiryDate: time.Date(2047, 1, 10, 15, 20, 52, 561387237, time.UTC),
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Owner{}, &models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Owner{}, &models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for _, owner := range owners {
		err = db.Debug().Model(&models.Owner{}).Create(&owner).Error
		if err != nil {
			log.Fatalf("cannot seed owner table: %v", err)
		}
	}

	for _, vehicle := range vehicles {
		err = db.Debug().Model(&models.Vehicle{}).Create(&vehicle).Error
		if err != nil {
			log.Fatalf("cannot seed vehicle table: %v", err)
		}
	}
}
