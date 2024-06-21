package controller

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"
	"tokoku_project/internal/model"
)

type CustomerController struct {
	model *model.CustomerModel
}

func NewCustomerController(m *model.CustomerModel) *CustomerController {
	return &CustomerController{
		model: m,
	}
}

func (cc *CustomerController) Register(userID uint) (model.Customer, error) {
	var newData model.Customer
	scanner := bufio.NewScanner(os.Stdin)
	newData.CreatedBy = userID
	fmt.Println("----------------------------------")
	fmt.Println("Tambah Pegawai")
	fmt.Println("----------------------------------")
	//----------------------------------START TAMPIL CUSTOMER-------------------------------------------
	DataCustomer, err := cc.model.GetCustomer()

	if err != nil {
		fmt.Println("Terjadi ERROR")
	} else {

		fmt.Println("\n--------------")
		fmt.Println("Daftar Customer")
		fmt.Println("--------------")

		w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
		//fmt.Fprintln(w, "----------------------------------------------------")
		fmt.Fprintln(w, "| id\t| Nama Customer\t|")
		fmt.Fprintln(w, "+\t+\t+")
		for _, Customers := range DataCustomer {
			fmt.Fprintln(w, "|", Customers.ID, "\t|", Customers.Nama, "\t|")
		}
		w.Flush()
	}
	//----------------------------------END TAMPIL CUSTOMER-------------------------------------------
	fmt.Print("\nMasukkan Nama Customer : ")
	scanner.Scan()
	newData.Nama = scanner.Text()
	fmt.Print("Masukkan Email Customer : ")
	scanner.Scan()
	newData.Email = scanner.Text()
	fmt.Print("Masukkan Nomor Handphone Customer : ")
	scanner.Scan()
	newData.Phone = scanner.Text()
	result, err := cc.model.Register(newData)
	if err != nil && !result {
		return model.Customer{}, errors.New("terjadi kesalahan saat melakukan register")
	}
	fmt.Print("\033[H\033[2J") //cls
	fmt.Println(newData.Nama, "berhasil ditambahkan ke Daftar Customer")
	return newData, nil
}
