package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"tokoku_project/internal/model"
)

type RestockController struct {
	model *model.RestockModel
}

func NewRestockController(m *model.RestockModel) *RestockController {
	return &RestockController{
		model: m,
	}
}

func (rc *RestockController) RestockBarang(userID uint) (model.Restock, error) {
	var newData model.Restock
	var confirm int
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\nMasukkan Nama Restock Barang : ")
		scanner.Scan() 
		newData.RestockBarang = scanner.Text()

		fmt.Print("Masukkan Jenis Restock Barang : ")
		scanner.Scan() 
		newData.RestockJenisBarang = scanner.Text()

		fmt.Print("Masukkan Harga Barang : ")
		fmt.Scanln(&newData.Harga)

		fmt.Println("\nNama Barang : ", newData.RestockBarang)
		fmt.Println("Jenis Barang : ", newData.RestockJenisBarang)
		fmt.Println("Harga Barang : ", newData.Harga)
		for {
			fmt.Println("Apakah data restock yang ingin ditambahkan sudah benar ?\n[1] YA\n[2] EDIT ULANG DATA RESTOCK\n[3] BATAL RESTOCK BARANG")
			fmt.Println("Input anda : ")
			fmt.Scanln(&confirm)
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
			fmt.Println("Restock Batal Ditambahkan")
			return model.Restock{}, nil
		}
	}

	newData.CreatedBy = userID

	result, err := rc.model.RestockBarang(newData)
	if err != nil && !result {
		return model.Restock{}, errors.New("terjadi masalah ketika menambahkan barang baru")
	}

	fmt.Print("\033[H\033[2J") //cls

	fmt.Print(newData.RestockBarang, " Restock berhasil ditambahkan ke Daftar Barang")
	return newData, nil
}