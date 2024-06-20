package controller

import (
	"fmt"
	"tokoku_project/internal/model"
)

type TransaksiController struct {
	model *model.TransaksiModel
}

func NewTransaksiController(m *model.TransaksiModel) *TransaksiController {
	return &TransaksiController{
		model: m,
	}
}

func (tc *TransaksiController) RestockBarang(idUser uint) {

	var newData model.Barang
	var newDataTrx model.Transaksi
	var newDataDTrx model.DetailTransaksi

	var idInput, TambahBarang int
	fmt.Println("-----------------")
	fmt.Println("Restock Barang")
	fmt.Println("-----------------")

	for {
		var temp string
		fmt.Print("\nMasukkan ID barang : ")
		_, err := fmt.Scanln(&idInput)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan input yang valid")
			continue
		} else {
			break
		}

	}

	newData, err := tc.model.GetSatuBarang(idInput)

	if err != nil {
		fmt.Println("Id yang anda masukkan tidak ada")
		return
	}

	fmt.Println("ID barang ditemukan dengan data sebagai berikut :")
	fmt.Println("\nNama Barang : ", newData.NamaBarang)
	fmt.Println("Jenis Barang : ", newData.JenisBarang)
	fmt.Println("Harga Barang : ", newData.Harga)
	fmt.Println("Stock Saat Ini : ", newData.Stock)

	for {
		var temp string
		fmt.Print("\nJumlah Tambahan Barang : ")
		_, err := fmt.Scanln(&TambahBarang)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan angka yang valid")
			continue
		} else {
			break
		}
	}

	var confirm int
	for {
		fmt.Println("\nApakah jumlah penambahan barang sudah benar ?\n[1] YA\n[2] BATAL RESTOCK BARANG")
		fmt.Print("\nInput anda : ")
		_, err := fmt.Scanln(&confirm)
		if confirm > 0 || confirm < 3 {
			break
		} else if err != nil {
			var temp string
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan input yang valid")
			continue
		} else {
			fmt.Println("- Masukkan input yang valid")
			continue
		}
	}
	if confirm == 1 {
		newData.Stock += uint(TambahBarang)

		newData, err := tc.model.UpdateInfoBarang(newData)
		if err != nil {
			fmt.Println("terjadi masalah ketika update Stock Barang")
			return
		}

		newDataTrx.IdUser = idUser
		newDataTrx.JenisTransaksi = "Restock"
		newDataTrx, err := tc.model.UpdateTransaksi(newDataTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update data Restock")
			return
		}

		newDataDTrx.IdTransaksi = newDataTrx.ID
		newDataDTrx.IdBarang = uint(idInput)
		newDataDTrx.Qty = uint(TambahBarang)
		newDataDTrx.TotalHarga = newData.Harga * uint(TambahBarang)

		newDataDTrx, err := tc.model.UpdateDetailTransaksi(newDataDTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update detail data Restock")
			return
		}

		newDataTrx.TotalQty = newDataDTrx.Qty
		newDataTrx.TotalHarga = newDataDTrx.TotalHarga

		newDataTrx, err = tc.model.UpdateTransaksi(newDataTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update data Restock")
			return
		}

		fmt.Print("\033[H\033[2J") //cls

		fmt.Println(newData.NamaBarang, "berhasil Restock sebanyak", TambahBarang, "unit")

	} else {
		fmt.Print("\033[H\033[2J") //cls
		fmt.Println("Restock barang dibatalkan")
		return
	}
}
