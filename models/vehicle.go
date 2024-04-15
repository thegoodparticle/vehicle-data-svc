package models

import "time"

type Vehicle struct {
	RegistrationNumber string    `json:"registrationNumber" gorm:"primary_key"`
	VehicleModel       string    `json:"model" gorm:"size:255;not null"`
	Company            string    `json:"company" gorm:"size:255;not null"`
	EngineNumber       string    `json:"engineNumber" gorm:"size:255;not null"`
	ChassisNumber      string    `json:"chassisNumber" gorm:"size:255;not null"`
	OwnerLicense       string    `json:"ownerLicenseNumber" gorm:"size:255;not null"`
	RegistrationDate   time.Time `json:"registrationDate" gorm:"not null"`
}
