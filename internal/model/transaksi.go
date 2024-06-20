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

type DetailTransaksiModel struct {
	db *gorm.DB
}

func NewTransaksiModel(connection *gorm.DB) *TransaksiModel {
	return &TransaksiModel{
		db: connection,
	}
}

func NewDetailTransaksiModel(connection *gorm.DB) *DetailTransaksiModel {
	return &DetailTransaksiModel{
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

func (dtm *DetailTransaksiModel) UpdateDetailTransaksi(newDetailTransaksi DetailTransaksi) (DetailTransaksi, error) {

	err := dtm.db.Save(&newDetailTransaksi).Error
	if err != nil {
		return DetailTransaksi{}, err
	}
	return newDetailTransaksi, nil
}
