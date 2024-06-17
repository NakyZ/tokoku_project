package model

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	NamaBarang  string `gorm:"type:varchar(50)"`
	JenisBarang string `gorm:"type:varchar(50)"`
	Harga       uint
	Stock       uint
	CreatedBy   uint
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type BarangModel struct {
	db *gorm.DB
}

func NewBarangModel(connection *gorm.DB) *BarangModel {
	return &BarangModel{
		db: connection,
	}
}

func (bm *BarangModel) TambahBarang(newBarang Barang) (bool, error) {
	err := bm.db.Create(&newBarang).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
