package models

import "time"

type Vehicle struct {
	RegistrationNumber string    `json:"registrationNumber,omitempty" gorm:"primary_key"`
	VehicleModel       string    `json:"model,omitempty" gorm:"size:255;not null"`
	Company            string    `json:"company,omitempty" gorm:"size:255;not null"`
	EngineNumber       string    `json:"engineNumber,omitempty" gorm:"size:255;not null"`
	ChassisNumber      string    `json:"chassisNumber,omitempty" gorm:"size:255;not null"`
	OwnerLicense       string    `json:"ownerLicenseNumber,omitempty" gorm:"size:255;not null"`
	RegistrationDate   time.Time `json:"registrationDate,omitempty" gorm:"not null"`
}
