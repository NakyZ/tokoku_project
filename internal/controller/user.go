package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
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

	fmt.Print("Login Berhasil")
	return result, nil
}

func (uc *UserController) Register() (model.User, error) {
	var newData model.User
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan username Pegawai	:")
	scanner.Scan()
	newData.Username = scanner.Text()
	fmt.Print("Masukkan Password	:")
	scanner.Scan()
	newData.Password = scanner.Text()
	fmt.Print("Masukkan Nama pegawai	:")
	scanner.Scan()
	newData.Nama = scanner.Text()
	newData.IsAdmin = false
	result, err := uc.model.Register(newData)
	if err != nil && !result {
		return model.User{}, errors.New("terjadi kesalahan saat melakukan register")
	}
	return newData, nil
}
