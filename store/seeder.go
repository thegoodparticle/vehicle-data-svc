package store

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/thegoodparticle/vehicle-data-svc/models"
)

var drivers = []models.Driver{
	{
		LicenseNumber: "KA20123456",
		DriverName:    "John Doe",
		LicenseDate:   time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
	},
	{
		LicenseNumber: "KA20231211",
		DriverName:    "Tim David",
		LicenseDate:   time.Date(2014, 1, 12, 15, 20, 52, 561387237, time.UTC),
	},
}

var vehicles = []models.Vehicle{
	{
		RegistrationNumber: "KA20AB1234",
		VehicleModel:       "Bolero",
		Company:            "Mahindra",
		EngineNumber:       "52WVC10338",
		ChassisNumber:      "QWERTY12321233ABC",
		OwnerLicense:       "KA20123456",
		RegistrationDate:   time.Date(2010, 5, 17, 20, 34, 58, 651387237, time.UTC),
	},
	{
		RegistrationNumber: "KA20CD5678",
		VehicleModel:       "Polo",
		Company:            "VW",
		EngineNumber:       "62WVC10338",
		ChassisNumber:      "ASDFGH12321233ABC",
		OwnerLicense:       "KA20231211",
		RegistrationDate:   time.Date(2015, 1, 12, 15, 20, 52, 561387237, time.UTC),
	},
}

func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Driver{}, &models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.Driver{}, &models.Vehicle{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for _, driver := range drivers {
		err = db.Debug().Model(&models.Driver{}).Create(&driver).Error
		if err != nil {
			log.Fatalf("cannot seed driver table: %v", err)
		}
	}

	for _, vehicle := range vehicles {
		err = db.Debug().Model(&models.Vehicle{}).Create(&vehicle).Error
		if err != nil {
			log.Fatalf("cannot seed vehicle table: %v", err)
		}
	}
}
