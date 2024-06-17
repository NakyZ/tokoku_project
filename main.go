package main

import (
	"fmt"
	"tokoku_project/config"

	"tokoku_project/internal/controller"
	"tokoku_project/internal/model"
)

func main() {
	fmt.Println("----------------------------------")
	fmt.Println("Selamat Datang di Aplikasi TOKOKU")
	fmt.Println("----------------------------------")
	setup := config.ImportSetting()
	connection, err := config.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database" /*, err.Error()*/)
		return
	}

	connection.AutoMigrate(&model.User{}, &model.Barang{})

	var inputMenu int

	um := model.NewUserModel(connection)
	uc := controller.NewUserController(um)

	bm := model.NewBarangModel(connection)
	bc := controller.NewBarangController(bm)

	// tu := models.NewTodoModel(connection)
	// tc := controllers.NewTodoController(tu)

	for inputMenu != 9 {
		var temp string
		fmt.Println("\nPilih menu")
		fmt.Println("\n1. Login")
		fmt.Println("9. Keluar")
		fmt.Print("\nMasukkan input: ")
		_, err := fmt.Scanln(&inputMenu)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("Input salah, silahkan coba lagi")
			continue
		}
		if inputMenu == 1 {

			data, err := uc.Login()
			if err != nil {
				//fmt.Println("Terjadi error pada saat login, error: ", err.Error())
				continue
			}
			var isLogin = true
			var inputMenu2 int

			for isLogin && data.IsAdmin {
				fmt.Println("\nSelamat datang", data.Nama, ",")
				fmt.Println("\nPilih menu")
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. Edit Informasi Barang")
				fmt.Println("3. Restock Barang")
				fmt.Println("4. Pembelian")
				fmt.Println("5. Tambah Pegawai")
				fmt.Println("6. Tambah Customer")
				fmt.Println("7. Kurangi Stok (Opsional)")
				fmt.Println("9. Keluar")
				fmt.Print("\nMasukkan input: ")
				fmt.Scanln(&inputMenu2)

				if inputMenu2 == 1 {
					bc.TambahBarang(data.ID)
				}
				if inputMenu2 == 9 {
					isLogin = false
				}
				if inputMenu2 != 1 || inputMenu2 != 9 {
					fmt.Print("\033[H\033[2J") //cls
					fmt.Print("Input anda salah atau fitur yang dipilih belum tersedia")
				}
				// } else if inputMenu2 == 1 {
				// 	_, err := tc.AddTodo(data.ID)
				// 	if err != nil {
				// 		fmt.Println("error ketika menambahkan aktivitas")
				// 		return
				// 	}
				// 	fmt.Println("berhasil menambahkan aktivitas")
				// }
			}
		}
		// } else if inputMenu == 2 {
		// 	uc.Register()
		// }
	}

	fmt.Println("\nterima kasih")
}
