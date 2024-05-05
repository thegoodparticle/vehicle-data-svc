package store

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/thegoodparticle/vehicle-data-svc/models"
)

type Store struct {
	DB *gorm.DB
}

func (s *Store) FindVehicleDataByRegistration(registrationNumber string) (*models.Vehicle, error) {
	var err error

	vehicle := &models.Vehicle{}
	err = s.DB.Debug().Model(&models.Vehicle{}).Where("registration_number = ?", registrationNumber).Take(vehicle).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return vehicle, nil
}

func (s *Store) FindVehicleDataByChassisEngine(engineNumber, chassisNumber string) (*models.Vehicle, error) {
	var err error

	vehicle := &models.Vehicle{}
	err = s.DB.Debug().Model(&models.Vehicle{}).Where("engine_number = ? and chassis_number = ?",
		engineNumber, chassisNumber).Take(vehicle).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return vehicle, nil
}

func (s *Store) FindOwnerDataByID(licenseNumber string) (*models.Owner, error) {
	var err error

	owner := &models.Owner{}
	err = s.DB.Debug().Model(&models.Owner{}).Where("owner_id = ?", licenseNumber).Take(owner).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return owner, nil
}
