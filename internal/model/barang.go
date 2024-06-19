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

func (bm *BarangModel) GetBarang() ([]Barang, error) {
	var result []Barang
	err := bm.db.Find(&result).Error
	if err != nil {
		return []Barang{}, err
	}
	return result, nil
}

func (bm *BarangModel) GetSatuBarang(id int) (Barang, error) {
	var result Barang
	err := bm.db.First(&result, id).Error
	if err != nil || bm.db.First(&result).RowsAffected == 0 {
		return Barang{}, err
	}
	return result, nil
}

func (bm *BarangModel) UpdateInfoBarang(newBarang Barang) (Barang, error) {

	err := bm.db.Save(&newBarang).Error
	if err != nil {
		return Barang{}, err
	}
	return newBarang, nil
}
