package controller

import (
	"errors"
	"fmt"
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
