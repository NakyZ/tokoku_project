package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Email string `gorm:"type:varchar(50)"`
	Nama string `gorm:"type:varchar(50)"`
	Phone    string `gorm:"type:varchar(50)"`
	CreatedBy uint
	// BirthDate time.Time `gorm:"type:date"`
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type CustomerModel struct{
	db *gorm.DB
}

func NewCustomerModel(connection *gorm.DB) *CustomerModel  {
	return &CustomerModel{
		db: connection, 
	}
}