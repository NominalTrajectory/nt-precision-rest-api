package main

import (
	"log"
	"os"

	"github.com/NominalTrajectory/nt-precision-rest-api/database"
	"github.com/NominalTrajectory/nt-precision-rest-api/server"
)

var (
	DBConnectionString string = os.Getenv("DB_CONNECTION_STRING")
	ListenAddress      string = os.Getenv("LISTEN_ADDRESS")
)

func main() {
	log.Println("Connecting to the database...")
	database.InitializeDatabase(DBConnectionString)
	log.Printf("Starting Precision API at %v", ListenAddress)
	server.InitializeServer(ListenAddress, DBConnectionString)
	err := server.Server.ListenAndServe()
	if err != nil {
		log.Fatalf("Precision API failed to start: %v", err)
	}
}
