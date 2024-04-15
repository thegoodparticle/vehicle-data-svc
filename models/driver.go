package models

import "time"

type Driver struct {
	LicenseNumber string    `json:"licenseNumber" gorm:"primary_key"`
	DriverName    string    `json:"name" gorm:"size:255;not null"`
	LicenseDate   time.Time `json:"licenseDate" gorm:"not null"`
}
