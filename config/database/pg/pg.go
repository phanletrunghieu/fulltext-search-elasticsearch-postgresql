package pg

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "123456"
	dbName     = "blog"
)

var db *gorm.DB

// Init .
func Init() func() {
	connectionString := fmt.Sprintf("host = %s port=%d user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	d, err := gorm.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = d
	return func() {
		err := db.Close()
		if err != nil {
			log.Println("Failed to close DB by error", err)
		}
	}
}

// GetDB .
func GetDB() *gorm.DB {
	return db
}
