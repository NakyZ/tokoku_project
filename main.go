package main

import (
	"fmt"
	"tokoku_project/config"
	"tokoku_project/internal/controller"
	"tokoku_project/internal/model"
)

func main() {
	// fmt.Println("----------------------------------")
	// fmt.Println("Selamat Datang di Aplikasi TOKOKU")
	// fmt.Println("----------------------------------")
	setup := config.ImportSetting()
	connection, err := config.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database" /*, err.Error()*/)
		return
	}

	connection.AutoMigrate(&model.User{}, &model.Barang{}, &model.Transaksi{}, &model.DetailTransaksi{})

	var inputMenu int

	um := model.NewUserModel(connection)
	uc := controller.NewUserController(um)

	bm := model.NewBarangModel(connection)
	bc := controller.NewBarangController(bm)

	tm := model.NewTransaksiModel(connection)
	tc := controller.NewTransaksiController(tm)

	dtm := model.NewDetailTransaksiModel(connection)
	dtc := controller.NewDetailTransaksiController(dtm)

	for inputMenu != 9 {
		var temp string
		fmt.Println("----------------------------------")
		fmt.Println("Selamat Datang di Aplikasi TOKOKU")
		fmt.Println("----------------------------------")
		fmt.Println("\nPilih menu")
		fmt.Println("\n1. Login")
		fmt.Println("9. Keluar")
		fmt.Print("\nMasukkan input: ")
		_, err := fmt.Scanln(&inputMenu)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Print("\033[H\033[2J") //cls
			fmt.Println("Input salah, silahkan coba lagi")
			continue
		}
		if inputMenu == 1 {

			data, err := uc.Login()
			if err != nil {
				fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				continue
			}
			var isLogin = true
			var inputMenu2 int
			var inputSubMenu int
			//------------------------------ LOGIN ADMIN -----------------------------------------//
			for isLogin && data.IsAdmin {
				fmt.Println("----------------------------------")
				fmt.Println("Menu Utama Admin")
				fmt.Println("----------------------------------")
				fmt.Println("Selamat datang", data.Nama, ",")
				fmt.Println("\nPilih menu")
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. Tampilkan Daftar Barang")
				fmt.Println("3. Edit Informasi Barang")
				fmt.Println("4. Restock Barang")
				fmt.Println("5. Pembelian")
				fmt.Println("6. Tambah Pegawai")
				fmt.Println("7. Tambah Customer")
				fmt.Println("8. Kurangi Stok (Opsional)")
				fmt.Println("\n99. Keluar")
				fmt.Print("\nMasukkan input: ")

				_, err := fmt.Scanln(&inputMenu2)
				if err != nil {
					fmt.Scanln(&temp)
					fmt.Print("\033[H\033[2J") //cls
					fmt.Println("Input salah, silahkan coba lagi")
					continue
				}

				switch inputMenu2 {
				case 1:
					bc.TambahBarang(data.ID)
				case 2:
					inputSubMenu = bc.GetBarang()
					if inputSubMenu == 1 {
						fmt.Println("Ini Fitur Restock tapi belum selesai dibuat")
					}
				case 3:
					fmt.Print("\033[H\033[2J") //cls
					bc.UpdateInfoBarang(data.ID)
				case 4:
					fmt.Print("\033[H\033[2J") //cls
					tc.RestockBarang(bc, dtc, data.ID)
				case 5:
					fmt.Print("\033[H\033[2J") //cls
					uc.Register()
				case 99:
					fmt.Print("\033[H\033[2J") //cls
					isLogin = false
				default:
					fmt.Print("\033[H\033[2J") //cls
					fmt.Println("Input anda salah atau fitur yang dipilih belum tersedia")
				}
			}
			//------------------------------ CLOSING LOGIN ADMIN -----------------------------------------//

			//------------------------------ LOGIN PEGAWAI -----------------------------------------//
			for isLogin && !data.IsAdmin {
				fmt.Println("----------------------------------")
				fmt.Println("Menu Utama Pegawai")
				fmt.Println("----------------------------------")
				fmt.Println("\nSelamat datang", data.Nama, ",")
				fmt.Println("\nPilih menu")
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. Edit Informasi Barang")
				fmt.Println("3. Restock Barang")
				fmt.Println("4. Pembelian")
				fmt.Println("6. Tambah Customer")
				fmt.Println("7. Kurangi Stok (Opsional)")
				fmt.Println("99. Keluar")
				fmt.Print("\nMasukkan input: ")
				fmt.Scanln(&inputMenu2)
				switch inputMenu2 {
				case 1:
					bc.TambahBarang(data.ID)
				case 2:
					bc.GetBarang()
				case 99:
					fmt.Print("\033[H\033[2J") //cls
					isLogin = false
				default:
					fmt.Print("\033[H\033[2J") //cls
					fmt.Print("Input anda salah atau fitur yang dipilih belum tersedia")
				}

			}
			//------------------------------ CLOSING LOGIN PEGAWAI -----------------------------------------//
		}

	}

	fmt.Println("\nterima kasih")
}
