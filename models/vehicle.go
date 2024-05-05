package models

import "time"

type Vehicle struct {
	RegistrationNumber     string    `json:"registrationNumber,omitempty" gorm:"primary_key"`
	VehicleModel           string    `json:"model,omitempty" gorm:"size:256;not null"`
	Company                string    `json:"company,omitempty" gorm:"size:256;not null"`
	EngineNumber           string    `json:"engineNumber,omitempty" gorm:"size:256;not null"`
	ChassisNumber          string    `json:"chassisNumber,omitempty" gorm:"size:256;not null"`
	OwnerID                string    `json:"ownerID,omitempty" gorm:"size:256;not null"`
	RegistrationIssueDate  time.Time `json:"registrationIssueDate,omitempty" gorm:"not null"`
	RegistrationExpiryDate time.Time `json:"registrationExpiryDate,omitempty" gorm:"not null"`
}
