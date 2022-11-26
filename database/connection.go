package database

import (
	"github.com/mattiadevivo/crm/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = connection
	if err := connection.AutoMigrate(&models.Customer{}); err != nil {
		return err
	}
	return nil
}
