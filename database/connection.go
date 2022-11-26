package database

import (
	"fmt"

	"github.com/mattiadevivo/crm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprint("Could not connect to the database", err))
	}

	DB = connection

	connection.AutoMigrate(&models.Customer{})
}
