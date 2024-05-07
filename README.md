protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative vehicledata/vehicle-data.proto

docker build -t vehicle-data-svc:v1.0.0 .

grpcurl -import-path . -proto vehicledata/vehicle-data.proto -plaintext -d '{"RegistrationNumber": "KA20AB1234"}' 127.0.0.1:8080 vehicledata.VehicleData.GetVehicleDataByRegistration

grpcurl -import-path . -proto vehicledata/vehicle-data.proto -plaintext -d '{"OwnerID": "KA20231211"}' 127.0.0.1:8080 vehicledata.VehicleData.GetOwnerDataByID



shebbar@Sandeshs-MacBook-Air vehicle-data-svc % kubectl create namespace vehicle
namespace/vehicle created

shebbar@Sandeshs-MacBook-Air vehicle-data-svc % kubectl --context=minikube apply -f ./deployment/postgres.yaml --namespace=vehicle                             
secret/postgres-secret created
persistentvolume/postgres-volume created
persistentvolumeclaim/postgres-volume-claim created
deployment.apps/postgres created
service/postgres created

shebbar@Sandeshs-MacBook-Air vehicle-data-svc % kubectl get pods -n=vehicle                                                                
NAME                        READY   STATUS    RESTARTS   AGE
postgres-575bb558dd-8xwrp   1/1     Running   0          4s

shebbar@Sandeshs-MacBook-Air vehicle-data-svc % kubectl exec -it postgres-575bb558dd-8xwrp -n vehicle -- psql -h localhost -U postgres --password -p 5432 vehicle_info
Password: 
psql (14.11 (Debian 14.11-1.pgdg120+2))
Type "help" for help.

vehicle_info=# \conninfo
You are connected to database "vehicle_info" as user "postgres" on host "localhost" (address "::1") at port "5432".

kubectl scale deployment --replicas=5 postgres

kubectl --context=minikube delete -f ./deployment/postgres.yaml --ignore-not-found --namespace=vehicle

eval  $(minikube docker-env)

docker build -t vehicle-data-svc:v1.1.0 .

shebbar@Sandeshs-MacBook-Air vehicle-data-svc % kubectl --context=minikube apply -f ./deployment/vehicle-data-svc.yaml --namespace=vehicle                    
secret/vehicle-data-svc-secret created
deployment.apps/vehicle-data-svc created
service/vehicle-data-svc created

shebbar@Sandeshs-MacBook-Air challan-data-svc % kubectl get pods -n vehicle
NAME                                READY   STATUS    RESTARTS   AGE
postgres-575bb558dd-8xwrp           1/1     Running   0          50m
vehicle-data-svc-6f4cb5bbb9-r6mml   1/1     Running   0          18m