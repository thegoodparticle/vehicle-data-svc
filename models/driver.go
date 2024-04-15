package models

import "time"

type Driver struct {
	LicenseNumber string    `json:"licenseNumber,omitempty" gorm:"primary_key"`
	DriverName    string    `json:"name,omitempty" gorm:"size:255;not null"`
	LicenseDate   time.Time `json:"licenseDate,omitempty" gorm:"not null"`
}
