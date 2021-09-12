package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/NominalTrajectory/nt-precision-rest-api/model"
	"github.com/julienschmidt/httprouter"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbaddr = flag.String("dbaddr", "", "The address of the database")
)

func main() {
	flag.Parse()

	db := setupDB(*dbaddr)

	router := httprouter.New()

	server := NewServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":5100", router))

}

func setupDB(dbAddress string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// Migrate to schema
	db.AutoMigrate(&model.Objective{})

	return db
}
