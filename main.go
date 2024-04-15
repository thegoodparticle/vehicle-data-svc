package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	grpchandler "github.com/thegoodparticle/vehicle-data-svc/grpc-handler"
	"github.com/thegoodparticle/vehicle-data-svc/models"
	"github.com/thegoodparticle/vehicle-data-svc/store"
	"google.golang.org/grpc"

	pb "github.com/thegoodparticle/vehicle-data-svc/vehicledata"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Printf("Error while loadinf env file - %s", err)
	}

	dbDriver, dbPort, dbHost := os.Getenv("DB_DRIVER"), os.Getenv("DB_PORT"), os.Getenv("DB_HOST")
	dbUser, dbPassword, dbName := os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")

	storeObj := store.Store{}

	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPassword)

	storeObj.DB, err = gorm.Open(dbDriver, dbURL)
	if err != nil {
		log.Printf("Cannot connect to %s database\n", dbDriver)
		log.Fatal("This is the error:", err)
	} else {
		log.Printf("We are connected to the %s database\n", dbDriver)
	}

	storeObj.DB.Debug().AutoMigrate(&models.Vehicle{}, &models.Driver{}) //database migration

	port := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVehicleDataServer(grpcServer, grpchandler.NewServer(&storeObj))
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Could not start grpc server. Error - %+v\n", err)
	}
}
