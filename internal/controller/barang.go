package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"tokoku_project/internal/model"
)

type BarangController struct {
	model *model.BarangModel
}

func NewBarangController(m *model.BarangModel) *BarangController {
	return &BarangController{
		model: m,
	}
}

func (bc *BarangController) TambahBarang(userID uint) (model.Barang, error) {
	var newData model.Barang

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("\nMasukkan Nama Barang : ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	newData.NamaBarang = scanner.Text()

	fmt.Print("Masukkan Jenis Barang : ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	newData.JenisBarang = scanner.Text()

	fmt.Print("Masukkan Harga Barang : ")
	fmt.Scanln(&newData.Harga)

	newData.CreatedBy = userID

	result, err := bc.model.TambahBarang(newData)
	if err != nil && !result {
		return model.Barang{}, errors.New("terjadi masalah ketika menambahkan barang baru")
	}

	fmt.Print("\033[H\033[2J") //cls

	fmt.Print(newData.NamaBarang, " berhasil ditambahkan ke Daftar Barang")
	return newData, nil
}
