package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"tokoku_project/internal/model"
)

type CustomerController struct {
	model *model.CustomerModel
}

func NewCustomerController(m *model.CustomerModel) *CustomerController  {
	return &CustomerController{
		model: m,
	}
}

func (cc *CustomerController) Register(userID uint) (model.Customer, error) {
	var newData model.Customer
	scanner := bufio.NewScanner(os.Stdin)
	newData.CreatedBy = userID

	fmt.Print("Masukkan Nama Customer	:")
	scanner.Scan()
	newData.Nama = scanner.Text()
	fmt.Print("Masukkan Email Customer	:")
	scanner.Scan()
	newData.Email = scanner.Text()
	fmt.Print("Masukkan Nomor Handphone Customer	:")
	scanner.Scan()
	newData.Phone = scanner.Text()
	result, err := cc.model.Register(newData)
	if err != nil && !result { 
		return model.Customer{}, errors.New("Terjadi kesalahan saat melakukan register")
	}
	return newData, nil
}