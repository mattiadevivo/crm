package database

import (
	"fmt"

	"github.com/mattiadevivo/crm/models"
)

var CustomerDb []models.Customer

func Migrate() {
	CustomerDb = []models.Customer{}
	for i := 0; i < 8; i++ {
		CustomerDb = append(CustomerDb, models.Customer{
			Id:        i,
			Name:      fmt.Sprintf("Customer%d", i),
			Role:      fmt.Sprintf("Role%d", i),
			Email:     fmt.Sprintf("Customer%d@gmail.com", i),
			Phone:     fmt.Sprintf("+39348112212%d", i),
			Contacted: false,
		})
	}
}
