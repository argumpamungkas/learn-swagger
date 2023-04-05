package repo

import (
	"DTS/Chapter-2/Sesi/sesi-5-swagger/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "gume98"
	port     = "5432"
	dbname   = "db-cars"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=disable", host, user, password, port, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	db.Debug().AutoMigrate(models.Car{})
	// db.Debug().Migrator().DropTable(models.Car{})
}

func GetDB() *gorm.DB {
	return db
}
