package grpchandler

import (
	"context"
	"errors"
	"log"

	"github.com/thegoodparticle/vehicle-data-svc/store"
	pb "github.com/thegoodparticle/vehicle-data-svc/vehicledata"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type vehicleDataServer struct {
	pb.UnimplementedVehicleDataServer
	dataStore *store.Store
}

func NewServer(dataStore *store.Store) *vehicleDataServer {
	s := &vehicleDataServer{dataStore: dataStore}
	return s
}

func (s *vehicleDataServer) GetVehicleDataByRegistration(ctx context.Context, registrationReq *pb.RegistrationRequest) (*pb.VehicleInfo, error) {
	if registrationReq == nil || registrationReq.RegistrationNumber == "" {
		return nil, errors.New("empty registration number in request")
	}

	vehicleData, err := s.dataStore.FindVehicleDataByRegistration(registrationReq.RegistrationNumber)
	if err != nil {
		log.Printf("failed to get vehicle data for registration %s. Error - %s", registrationReq.RegistrationNumber, err)
		return nil, errors.New("failed to get driver data. error " + err.Error())
	}

	return &pb.VehicleInfo{
		RegistrationNumber: registrationReq.RegistrationNumber,
		VehicleModel:       vehicleData.VehicleModel,
		Company:            vehicleData.Company,
		RegistrationDate:   timestamppb.New(vehicleData.RegistrationDate),
	}, nil
}

func (s *vehicleDataServer) GetVehicleDataByChassisEngine(ctx context.Context, chassisEngineReq *pb.ChassisEngineNumberRequest) (*pb.VehicleInfo, error) {
	if chassisEngineReq == nil || chassisEngineReq.ChassisNumber == "" || chassisEngineReq.EngineNumber == "" {
		return nil, errors.New("empty engine/chassis number in request")
	}

	vehicleData, err := s.dataStore.FindVehicleDataByChassisEngine(chassisEngineReq.EngineNumber, chassisEngineReq.ChassisNumber)
	if err != nil {
		log.Printf("failed to get vehicle data for engine/chassis %s/%s. Error - %s",
			chassisEngineReq.EngineNumber, chassisEngineReq.ChassisNumber, err)
		return nil, errors.New("failed to get driver data. error " + err.Error())
	}

	return &pb.VehicleInfo{
		RegistrationNumber: vehicleData.RegistrationNumber,
		VehicleModel:       vehicleData.VehicleModel,
		Company:            vehicleData.Company,
		RegistrationDate:   timestamppb.New(vehicleData.RegistrationDate),
	}, nil
}

func (s *vehicleDataServer) GetDriverDataByLicenseNumber(ctx context.Context, licenseReq *pb.LicenseRequest) (*pb.DriverInfo, error) {
	if licenseReq == nil || licenseReq.LicenseNumber == "" {
		return nil, errors.New("empty license number found in request")
	}

	driverData, err := s.dataStore.FindDriverDataByLicenseNumber(licenseReq.LicenseNumber)
	if err != nil {
		log.Printf("failed to get driver data for license %s. Error - %s", licenseReq.LicenseNumber, err)
		return nil, errors.New("failed to get driver data. error " + err.Error())
	}

	return &pb.DriverInfo{
		LicenseNumber: driverData.LicenseNumber,
		DriverName:    driverData.DriverName,
		LicenseDate:   timestamppb.New(driverData.LicenseDate),
	}, nil
}
