package models

type Customer struct {
	Id        int `gorm:"primaryKey"`
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}
