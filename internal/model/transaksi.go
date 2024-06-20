package model

import "gorm.io/gorm"

type Transaksi struct {
	gorm.Model
	IdUser           uint
	IdCustomer       *uint
	JenisTransaksi   string `gorm:"type:varchar(50)"`
	TotalHarga       uint
	TotalQty         uint
	DetailTransaksis []DetailTransaksi `gorm:"foreignKey:IdTransaksi"`
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

type DetailTransaksi struct {
	IdTransaksi uint   `gorm:"primaryKey"`
	IdBarang    uint   `gorm:"primaryKey"`
	Keterangan  string `gorm:"type:varchar(100)"`
	TotalHarga  uint
	Qty         uint
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}
type TransaksiModel struct {
	db *gorm.DB
}

func NewTransaksiModel(connection *gorm.DB) *TransaksiModel {
	return &TransaksiModel{
		db: connection,
	}
}

func (tm *TransaksiModel) UpdateTransaksi(newTransaksi Transaksi) (Transaksi, error) {

	err := tm.db.Save(&newTransaksi).Error
	if err != nil {
		return Transaksi{}, err
	}
	return newTransaksi, nil
}

func (tm *TransaksiModel) UpdateDetailTransaksi(newDetailTransaksi DetailTransaksi) (DetailTransaksi, error) {

	err := tm.db.Save(&newDetailTransaksi).Error
	if err != nil {
		return DetailTransaksi{}, err
	}
	return newDetailTransaksi, nil
}

func (tm *TransaksiModel) GetSatuBarang(id int) (Barang, error) {
	var result Barang
	err := tm.db.First(&result, id).Error
	if err != nil || tm.db.First(&result).RowsAffected == 0 {
		return Barang{}, err
	}
	return result, nil
}

func (tm *TransaksiModel) UpdateInfoBarang(newBarang Barang) (Barang, error) {

	err := tm.db.Save(&newBarang).Error
	if err != nil {
		return Barang{}, err
	}
	return newBarang, nil
}

func (tm *TransaksiModel) GetSatuCustomer(id int) (Customer, error) {
	var result Customer
	err := tm.db.First(&result, id).Error
	if err != nil || tm.db.First(&result).RowsAffected == 0 {
		return Customer{}, err
	}
	return result, nil
}

func (tm *TransaksiModel) GetBarang() ([]Barang, error) {
	var result []Barang
	err := tm.db.Find(&result).Error
	if err != nil {
		return []Barang{}, err
	}
	return result, nil
}
