package database

import (
	"fmt"

	"github.com/mattiadevivo/crm/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(dsn string) error {
	log.Debug().Msg(fmt.Sprintf("Trying to connect to db: %s", dsn))
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

func Migrate() error {
	if DB == nil {
		return fmt.Errorf("DB connection not set!")
	}
	if err := DB.AutoMigrate(&models.Customer{}); err != nil {
		return err
	}
	if tx := DB.First(&models.Customer{Id: 1}); tx.Error == nil {
		// At least one row present in db, so don't insert
		return nil
	}
	for i := 1; i <= 8; i++ {
		if tx := DB.Create(&models.Customer{
			Id:        i,
			Name:      fmt.Sprintf("Customer%d", i),
			Role:      fmt.Sprintf("Role%d", i),
			Email:     fmt.Sprintf("Customer%d@gmail.com", i),
			Phone:     fmt.Sprintf("+39348112212%d", i),
			Contacted: false,
		}); tx.Error != nil {
			return tx.Error
		}
	}
	return nil
}
