package main

import (
	"fmt"
	"tokoku_project/config"

	"tokoku_project/internal/controller"
	"tokoku_project/internal/model"
)

func main() {
	fmt.Println("Test Main Program")
	setup := config.ImportSetting()
	connection, err := config.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database" /*, err.Error()*/)
		return
	}

	connection.AutoMigrate(&model.User{})

	var inputMenu int

	um := model.NewUserModel(connection)
	uc := controller.NewUserController(um)

	// tu := models.NewTodoModel(connection)
	// tc := controllers.NewTodoController(tu)

	for inputMenu != 9 {
		fmt.Println("\nPilih menu")
		fmt.Println("1. Login")
		fmt.Println("9. Keluar")
		fmt.Print("Masukkan input: ")
		fmt.Scanln(&inputMenu)
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
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
				}
				if inputMenu2 == 6 {
					uc.Register()
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
			for isLogin && !data.IsAdmin {
				fmt.Println("\nSelamat datang", data.Nama, ",")
				fmt.Println("\nPilih menu")
				fmt.Println("1. Tambah Barang")
				fmt.Println("2. Edit Informasi Barang")
				fmt.Println("3. Restock Barang")
				fmt.Println("4. Pembelian")
				fmt.Println("6. Tambah Customer")
				fmt.Println("7. Kurangi Stok (Opsional)")
				fmt.Println("9. Keluar")
				fmt.Print("Masukkan input: ")
				fmt.Scanln(&inputMenu2)
				if inputMenu2 == 9 {
					isLogin = false
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
