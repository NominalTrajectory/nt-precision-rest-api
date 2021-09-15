package database

import (
	"fmt"

	"github.com/NominalTrajectory/nt-precision-rest-api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase(dbConnectionString string) {
	db, err := gorm.Open(postgres.Open(dbConnectionString))
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}
	// Migrate to schema
	db.AutoMigrate(&model.Objective{})

	DB = db
}
