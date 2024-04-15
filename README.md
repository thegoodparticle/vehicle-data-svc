protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative vehicledata/vehicle-data.proto

docker build -t vehicle-data-svc:v1.0.0 .

grpcurl -import-path . -proto vehicledata/vehicle-data.proto -plaintext -d '{"RegistrationNumber": "KA20AB1234"}' 127.0.0.1:8080 vehicledata.VehicleData.GetVehicleDataByRegistration

grpcurl -import-path . -proto vehicledata/vehicle-data.proto -plaintext -d '{"LicenseNumber": "KA20231211"}' 127.0.0.1:8080 vehicledata.VehicleData.GetDriverDataByLicenseNumber