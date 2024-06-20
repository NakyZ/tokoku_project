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
		fmt.Print("Masukkan Stock  : ")
		scanner.Scan() 
		fmt.Scanln(&newData.Restock)

		fmt.Println("Restock Barang : ", newData.Restock)

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

	newData.Restock = userID

	result, err := rc.model.RestockBarang(newData)
	if err != nil && !result {
		return model.Restock{}, errors.New("terjadi masalah ketika menambahkan barang baru")
	}

	fmt.Print("\033[H\033[2J") //cls

	fmt.Print(newData.Restock, " Restock berhasil ditambahkan ke Daftar Stock")
	return newData, nil
}