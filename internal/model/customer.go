package model

import "gorm.io/gorm"

type Customer struct {
	gorm.Model
	Email     string `gorm:"type:varchar(50)"`
	Nama      string `gorm:"type:varchar(50)"`
	Phone     string `gorm:"type:varchar(50)"`
	CreatedBy uint
	// BirthDate time.Time `gorm:"type:date"`
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type CustomerModel struct {
	db *gorm.DB
}

func NewCustomerModel(connection *gorm.DB) *CustomerModel {
	return &CustomerModel{
		db: connection,
	}
}

func (cm *CustomerModel) Register(NewCustomer Customer) (bool, error) {
	err := cm.db.Create(&NewCustomer).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (cm *CustomerModel) GetSatuCustomer(id int) (Customer, error) {
	var result Customer
	err := cm.db.First(&result, id).Error
	if err != nil || cm.db.First(&result).RowsAffected == 0 {
		return Customer{}, err
	}
	return result, nil
}
