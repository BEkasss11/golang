package initializers

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	dsn := "host=localhost user=postgres password=ADMIN port=5432 dbname=GoProject sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})	

	if err != nil {
		panic("Could not connect to DATABASE")
	}
	fmt.Println("Success")
}
