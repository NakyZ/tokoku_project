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
	var confirm int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nMasukkan Nama Barang : ")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		newData.NamaBarang = scanner.Text()

		fmt.Print("Masukkan Jenis Barang : ")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		newData.JenisBarang = scanner.Text()

		fmt.Print("Masukkan Harga Barang : ")
		fmt.Scanln(&newData.Harga)

		fmt.Print("Masukkan Stock Barang : ")
		fmt.Scanln(&newData.Stock)

		fmt.Println("\nNama Barang : ", newData.NamaBarang)
		fmt.Println("Jenis Barang : ", newData.JenisBarang)
		fmt.Println("Harga Barang : ", newData.Harga)
		fmt.Println("Stock Barang : ", newData.Stock)
		for {
			fmt.Println("\nApakah data barang yang ingin ditambahkan sudah benar ?\n[1] YA\n[2] EDIT ULANG DATA BARANG\n[3] BATAL TAMBAH BARANG")
			fmt.Println("Input anda : ")
			fmt.Scan(&confirm)
			if confirm > 0 || confirm < 4 {
				break
			} else {
				fmt.Println("Input Salah, coba lagi")
			}
		}
		if confirm == 1 {
			break
		} else if confirm == 2 {
			continue
		} else if confirm == 3 {
			fmt.Print("\033[H\033[2J") //cls
			fmt.Println("Barang Batal Ditambahkan")
			return model.Barang{}, nil
		}
	}

	newData.CreatedBy = userID

	result, err := bc.model.TambahBarang(newData)
	if err != nil && !result {
		return model.Barang{}, errors.New("terjadi masalah ketika menambahkan barang baru")
	}

	fmt.Print("\033[H\033[2J") //cls

	fmt.Print(newData.NamaBarang, " berhasil ditambahkan ke Daftar Barang")
	return newData, nil
}
