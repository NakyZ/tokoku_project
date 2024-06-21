package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"tokoku_project/internal/model"
)

type UserController struct {
	model *model.UserModel
}

func NewUserController(m *model.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Login() (model.User, error) {
	var username, password string

	fmt.Print("\033[H\033[2J") //cls

	fmt.Print("----------------------")
	fmt.Print("\nLogin\n")
	fmt.Print("----------------------")
	fmt.Print("\n\nMasukkan username : ")
	fmt.Scanln(&username)
	fmt.Print("Masukkan password : ")
	fmt.Scanln(&password)
	result, err := uc.model.Login(username, password)
	if err != nil {
		return model.User{}, errors.New("kombinasi username & password tidak ditemukan")
	}

	fmt.Print("\033[H\033[2J") //cls

	fmt.Println("Login Berhasil")
	return result, nil
}

func (uc *UserController) Register() (model.User, error) {
	var newData model.User
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\033[H\033[2J") //cls
	fmt.Println("----------------------------------")
	fmt.Println("Tambah Pegawai")
	fmt.Println("----------------------------------")

	//----------------------------------START TAMPIL CUSTOMER-------------------------------------------
	DataUser, err := uc.model.GetUser()

	if err != nil {
		fmt.Println("Terjadi ERROR")
	} else {

		fmt.Println("\n--------------")
		fmt.Println("Daftar User")
		fmt.Println("--------------")

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		//fmt.Fprintln(w, "----------------------------------------------------")
		fmt.Fprintln(w, "| id\t| Nama User\t|")
		fmt.Fprintln(w, "+\t+\t+")
		for _, Users := range DataUser {
			fmt.Fprintln(w, "|", Users.ID, "\t|", Users.Nama, "\t|")
		}
		w.Flush()
	}
	//----------------------------------END TAMPIL CUSTOMER-------------------------------------------
	fmt.Print("\nMasukkan username Pegawai Baru : ")
	scanner.Scan()
	newData.Username = scanner.Text()
	fmt.Print("Masukkan Password Baru: ")
	scanner.Scan()
	newData.Password = scanner.Text()
	fmt.Print("Masukkan Nama Pegawai Baru : ")
	scanner.Scan()
	newData.Nama = scanner.Text()
	newData.IsAdmin = false
	result, err := uc.model.Register(newData)
	if err != nil && !result {
		return model.User{}, errors.New("terjadi kesalahan saat melakukan register")
	}
	fmt.Print("\033[H\033[2J") //cls
	fmt.Println(newData.Nama, "berhasil ditambahkan ke Daftar Pegawai")
	return newData, nil
}
