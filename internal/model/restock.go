package model

import "gorm.io/gorm"

type Restock struct {
	gorm.Model
	RestockBarang  string `gorm:"type:varchar(50)"`
	RestockJenisBarang string `gorm:"type:varchar(50)"`
	Harga       uint
	Stock       uint
	CreatedBy   uint
}

type RestockModel struct {
	db *gorm.DB
}

func NewRestockModel(connection *gorm.DB) *RestockModel {
	return &RestockModel{
		db: connection,
	}
}

func (rm *RestockModel) RestockBarang(newRestock Restock) (bool, error) {
	err := rm.db.Create(&newRestock).Error
	if err != nil {
		return false, err
	}
	return true, nil
}