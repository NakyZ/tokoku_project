package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
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
	fmt.Print("\033[H\033[2J") //cls
	for {

		fmt.Println("----------------------------------")
		fmt.Println("Tambah Barang")
		fmt.Println("----------------------------------")
		fmt.Print("\nMasukkan Nama Barang : ")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		newData.NamaBarang = scanner.Text()

		fmt.Print("Masukkan Jenis Barang : ")
		scanner.Scan() // use `for scanner.Scan()` to keep reading
		newData.JenisBarang = scanner.Text()

		for {
			var temp string
			fmt.Print("Masukkan Harga Barang : ")
			_, err := fmt.Scanln(&newData.Harga)
			if err != nil {
				fmt.Scanln(&temp)
				fmt.Println("-Masukkan angka yang valid")
				continue
			} else {
				break
			}
		}

		fmt.Println("\nNama Barang : ", newData.NamaBarang)
		fmt.Println("Jenis Barang : ", newData.JenisBarang)
		fmt.Println("Harga Barang : ", newData.Harga)
		for {
			fmt.Println("\nApakah data barang yang ingin ditambahkan sudah benar ?\n[1] YA\n[2] EDIT ULANG DATA BARANG\n[3] BATAL TAMBAH BARANG")
			fmt.Print("\nInput anda : ")
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
			fmt.Print("\033[H\033[2J") //cls
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

	fmt.Println(newData.NamaBarang, " berhasil ditambahkan ke Daftar Barang")
	return newData, nil
}

func (bc *BarangController) GetBarang() int {
	var nextMenu int

	result, err := bc.model.GetBarang()

	if err != nil {
		fmt.Println("Terjadi ERROR")
	} else {
		fmt.Print("\033[H\033[2J") //cls

		fmt.Println("--------------")
		fmt.Println("Daftar Barang")
		fmt.Println("--------------")

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		//fmt.Fprintln(w, "----------------------------------------------------")
		fmt.Fprintln(w, "| id\t| Nama Barang\t| Jenis Barang\t| Harga\t| Stok\t|\t")
		fmt.Fprintln(w, "+\t+\t+\t+\t+\t+\t")
		for _, databarang := range result {
			fmt.Fprintln(w, "|", databarang.ID, "\t|", databarang.NamaBarang, "\t|", databarang.JenisBarang, "\t|", databarang.Harga, "\t|", databarang.Stock, "\t|\t")
		}
		w.Flush()
	}

	fmt.Println("\nPilih Menu :")
	fmt.Println("[1] RESTOCK BARANG")
	fmt.Println("[99] KEMBALI KE MENU UTAMA")
	for {
		var temp string
		fmt.Print("\nMasukkan Input Anda : ")
		_, err := fmt.Scanln(&nextMenu)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan input yang valid")
			continue
		}
		switch nextMenu {
		case 1:
			fmt.Print("\033[H\033[2J") //cls
			return nextMenu
		case 99:
			fmt.Print("\033[H\033[2J") //cls
			return nextMenu
		default:
			fmt.Println("- Masukkan input yang valid")
		}

	}

}

func (bc *BarangController) UpdateInfoBarang(idUser uint) {

	var newData model.Barang
	var idInput int
	fmt.Println("-----------------")
	fmt.Println("Edit Data Barang")
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

	newData, err := bc.model.GetSatuBarang(idInput)

	if err != nil {
		fmt.Println("Id yang anda masukkan tidak ada")
		return
	}

	fmt.Println("ID barang ditemukan dengan data sebagai berikut :")
	fmt.Println("\nNama Barang : ", newData.NamaBarang)
	fmt.Println("Jenis Barang : ", newData.JenisBarang)
	fmt.Println("Harga Barang : ", newData.Harga)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nMasukkan Nama Baru Barang : ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	newData.NamaBarang = scanner.Text()

	fmt.Print("Masukkan Jenis Baru Barang : ")
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	newData.JenisBarang = scanner.Text()

	for {
		var temp string
		fmt.Print("Masukkan Harga Baru Barang : ")
		_, err := fmt.Scanln(&newData.Harga)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan angka yang valid")
			continue
		} else {
			break
		}
	}

	fmt.Println("\nNama Baru Barang : ", newData.NamaBarang)
	fmt.Println("Jenis Baru Barang : ", newData.JenisBarang)
	fmt.Println("Harga Baru Barang : ", newData.Harga)

	var confirm int
	for {
		fmt.Println("\nApakah data barang baru barang sudah benar ?\n[1] YA\n[2] BATAL UPDATE BARANG")
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
		newData.CreatedBy = idUser
		newData, err := bc.model.UpdateInfoBarang(newData)
		if err != nil {
			fmt.Println("terjadi masalah ketika update data barang")
			return
		}

		fmt.Print("\033[H\033[2J") //cls

		fmt.Println(newData.NamaBarang, " berhasil ditambahkan ke Daftar Barang")

	} else {
		fmt.Print("\033[H\033[2J") //cls
		fmt.Println("Edit data barang dibatalkan")
		return
	}

}
