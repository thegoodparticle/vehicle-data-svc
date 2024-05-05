package models

import "time"

type Owner struct {
	OwnerID      string    `json:"ownerID,omitempty" gorm:"primary_key"`
	FirstName    string    `json:"firstName,omitempty" gorm:"size:256;not null"`
	LastName     string    `json:"lastName,omitempty" gorm:"size:256;not null"`
	MobileNumber string    `json:"mobileNumber,omitempty" gorm:"size:10;not null"`
	AadharNumber string    `json:"aaharNumber,omitempty" gorm:"size:12;not null"`
	Address      string    `json:"address,omitempty" gorm:"size:1024;not null"`
	CreatedAt    time.Time `json:"createdDate,omitempty" gorm:"not null"`
}
