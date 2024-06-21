package controller

import (
	"fmt"
	"os"
	"text/tabwriter"
	"tokoku_project/internal/model"
)

type Keranjang struct {
	BarangDibeli []model.Barang
	JumlahBeli   []int
	TotalHarga   int
}

type TransaksiController struct {
	model *model.TransaksiModel
}

func NewTransaksiController(m *model.TransaksiModel) *TransaksiController {
	return &TransaksiController{
		model: m,
	}
}

func (tc *TransaksiController) RestockBarang(idUser uint) {

	var newData model.Barang
	var newDataTrx model.Transaksi
	var newDataDTrx model.DetailTransaksi

	var idInput, JumlahBeli int
	fmt.Println("-----------------")
	fmt.Println("Restock Barang")
	fmt.Println("-----------------")
	//----------------------------------START TAMPIL BARANG-------------------------------------------
	result, err := tc.model.GetBarang()

	if err != nil {
		fmt.Println("Terjadi ERROR")
	} else {

		fmt.Println("\n--------------")
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
	//----------------------------------END TAMPIL BARANG-------------------------------------------

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

	newData, err = tc.model.GetSatuBarang(idInput)

	if err != nil {
		fmt.Println("Id yang anda masukkan tidak ada")
		return
	}

	fmt.Println("ID barang ditemukan dengan data sebagai berikut :")
	fmt.Println("\nNama Barang : ", newData.NamaBarang)
	fmt.Println("Jenis Barang : ", newData.JenisBarang)
	fmt.Println("Harga Barang : ", newData.Harga)
	fmt.Println("Stock Saat Ini : ", newData.Stock)

	for {
		var temp string
		fmt.Print("\nJumlah Tambahan Barang : ")
		_, err := fmt.Scanln(&JumlahBeli)
		if err != nil {
			fmt.Scanln(&temp)
			fmt.Println("- Masukkan angka yang valid")
			continue
		} else {
			break
		}
	}

	var confirm int
	for {
		fmt.Println("\nApakah jumlah penambahan barang sudah benar ?\n[1] YA\n[2] BATAL RESTOCK BARANG")
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
		newData.Stock += uint(JumlahBeli)

		newData, err := tc.model.UpdateInfoBarang(newData)
		if err != nil {
			fmt.Println("terjadi masalah ketika update Stock Barang")
			return
		}

		newDataTrx.IdUser = idUser
		newDataTrx.JenisTransaksi = "Restock"
		newDataTrx, err := tc.model.UpdateTransaksi(newDataTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update data Restock")
			return
		}

		newDataDTrx.IdTransaksi = newDataTrx.ID
		newDataDTrx.IdBarang = uint(idInput)
		newDataDTrx.Qty = uint(JumlahBeli)
		newDataDTrx.TotalHarga = newData.Harga * uint(JumlahBeli)

		newDataDTrx, err := tc.model.UpdateDetailTransaksi(newDataDTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update detail data Restock")
			return
		}

		newDataTrx.TotalQty = newDataDTrx.Qty
		newDataTrx.TotalHarga = newDataDTrx.TotalHarga

		newDataTrx, err = tc.model.UpdateTransaksi(newDataTrx)
		if err != nil {
			fmt.Println("terjadi masalah ketika update data Restock")
			return
		}

		fmt.Print("\033[H\033[2J") //cls

		fmt.Println(newData.NamaBarang, "berhasil Restock sebanyak", JumlahBeli, "unit")

	} else {
		fmt.Print("\033[H\033[2J") //cls
		fmt.Println("Restock barang dibatalkan")
		return
	}
}

func (tc *TransaksiController) Pembelian(DataUser model.User) {

	var Keranjang Keranjang
	var DataBarang model.Barang
	var DataCustomer model.Customer
	var DataTRX model.Transaksi
	var DetailTRX model.DetailTransaksi
	var err error
	var idInput, JumlahBeli int
	fmt.Println("-----------------")
	fmt.Println("Pembelian")
	fmt.Println("-----------------")
	//----------------------------------START TAMPIL CUSTOMER-------------------------------------------
	result, err := tc.model.GetCustomer()

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
		for _, Customers := range result {
			fmt.Fprintln(w, "|", Customers.ID, "\t|", Customers.Nama, "\t|")
		}
		w.Flush()
	}
	//----------------------------------END TAMPIL CUSTOMER-------------------------------------------

	//GET DATA Customer ---------------------------------------START-----------------------------------------
	for {
		for {
			var temp string
			fmt.Print("\nMasukkan ID Customer : ")
			_, err := fmt.Scanln(&idInput)
			if err != nil {
				fmt.Scanln(&temp)
				fmt.Println("- Masukkan input yang valid")
				continue
			} else {
				break
			}

		}

		DataCustomer, err = tc.model.GetSatuCustomer(idInput)

		if err != nil {
			fmt.Print("\033[H\033[2J") //cls
			fmt.Println("Id Customer yang anda masukkan tidak ada")
			continue
		}
		break
	}
	// fmt.Println("ID Customer ditemukan dengan data sebagai berikut :")
	// fmt.Println("\nNama Customer : ", DataCustomer.Nama)
	// fmt.Println("Email Customer : ", DataCustomer.Email)
	// fmt.Println("No HP Customer : ", DataCustomer.Phone)
	//GET DATA CUSTOMER -------------------------------------END-------------------------------------------
	for {
		fmt.Print("\033[H\033[2J") //cls

		fmt.Println("Halo", DataCustomer.Nama, ", belanja apa hari ini ?")

		if Keranjang.TotalHarga > 0 {
			fmt.Println("\nBarang di keranjang kamu sekarang :")

			w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
			for i, BarangKeranjang := range Keranjang.BarangDibeli {
				//fmt.Fprintln(w, "----------------------------------------------------")
				fmt.Fprintln(w, "-", BarangKeranjang.NamaBarang, "\t- Jumlah :", Keranjang.JumlahBeli[i], "\t- Total Harga :", uint(Keranjang.JumlahBeli[i])*BarangKeranjang.Harga)
			}
			w.Flush()
			fmt.Println("\nTotal Semua barang :", Keranjang.TotalHarga)
		}

		//----------------------------------START TAMPIL BARANG-------------------------------------------
		result, err := tc.model.GetBarang()

		if err != nil {
			fmt.Println("Terjadi ERROR")
		} else {

			fmt.Println("\n--------------")
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
		//----------------------------------END TAMPIL BARANG-------------------------------------------

		//GET DATA BARANG ---------------------------------------START-----------------------------------------
		for {
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

			DataBarang, err = tc.model.GetSatuBarang(idInput)

			if err != nil {
				fmt.Println("Id Barang yang anda masukkan tidak ada")
				continue
			}
			break
		}
		fmt.Println("ID barang ditemukan dengan data sebagai berikut :")
		fmt.Println("\nNama Barang : ", DataBarang.NamaBarang)
		fmt.Println("Jenis Barang : ", DataBarang.JenisBarang)
		fmt.Println("Harga Barang : ", DataBarang.Harga)
		fmt.Println("Stock Saat Ini : ", DataBarang.Stock)
		//GET DATA BARANG -------------------------------------END-------------------------------------------
		//GET Input Jumlah Barang ----------------------------------START-------------------------------------------
		for {
			var temp string
			fmt.Print("\nJumlah Pembelian Barang : ")
			_, err := fmt.Scanln(&JumlahBeli)
			if err != nil {
				fmt.Scanln(&temp)
				fmt.Println("- Masukkan angka yang valid")
				continue
			} else if JumlahBeli > int(DataBarang.Stock) {
				fmt.Println("Jumlah pembelian melebihi stok yang tersedia, Harap input ulang")
				continue
			} else {
				var BarangLain bool = true
				for i, BarangKeranjang := range Keranjang.BarangDibeli {
					if DataBarang.ID == BarangKeranjang.ID {
						if Keranjang.JumlahBeli[i]+JumlahBeli > int(BarangKeranjang.Stock) {
							fmt.Println("Jumlah pembelian melebihi stok yang tersedia, Harap input ulang")
							BarangLain = false
							break
						} else if Keranjang.JumlahBeli[i]+JumlahBeli <= 0 {
							Keranjang.TotalHarga -= Keranjang.JumlahBeli[i] * int(BarangKeranjang.Harga)
							Keranjang.BarangDibeli = append(Keranjang.BarangDibeli[:i], Keranjang.BarangDibeli[i+1:]...)
							Keranjang.JumlahBeli = append(Keranjang.JumlahBeli[:i], Keranjang.JumlahBeli[i+1:]...)
							BarangLain = false
							fmt.Println("Barang berhasil dikeluarkan dari keranjang")
							break
						}
						Keranjang.JumlahBeli[i] += JumlahBeli
						Keranjang.TotalHarga += int(DataBarang.Harga) * JumlahBeli
						fmt.Println("Barang di keranjang berhasil diperbarui")
						BarangLain = false
						break
					}
				}
				if BarangLain && JumlahBeli > 0 {
					Keranjang.BarangDibeli = append(Keranjang.BarangDibeli, DataBarang)
					Keranjang.JumlahBeli = append(Keranjang.JumlahBeli, JumlahBeli)
					Keranjang.TotalHarga += int(DataBarang.Harga) * JumlahBeli
					fmt.Println("Barang berhasil ditambah ke keranjang")
					break
				} else if !BarangLain {
					break
				} else {
					fmt.Println("- Masukkan angka yang valid")
					continue
				}
			}
		}

		var confirm int
		for {
			fmt.Println("\nApakah ingin menambahkan barang lain ?\n[1] YA\n[2] TIDAK\n[3] BATAL BELI SEMUA BARANG")
			fmt.Print("\nInput anda : ")
			_, err := fmt.Scanln(&confirm)
			if confirm > 0 && confirm < 4 {
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
		if confirm == 2 {
			break
		} else if confirm == 3 {
			return
		}

	}

	//GET Input Jumlah Barang ----------------------------------End-------------------------------------------

	DataTRX.IdUser = DataUser.ID
	DataTRX.IdCustomer = &DataCustomer.ID
	DataTRX.JenisTransaksi = "Pembelian"
	DataTRX, err = tc.model.UpdateTransaksi(DataTRX)
	if err != nil {
		fmt.Println("terjadi masalah ketika update data Pembelian")
		return
	}

	for i, BarangTerjual := range Keranjang.BarangDibeli {

		BarangTerjual.Stock -= uint(Keranjang.JumlahBeli[i])
		BarangTerjual, err = tc.model.UpdateInfoBarang(BarangTerjual)
		if err != nil {
			fmt.Println("terjadi masalah ketika update Pembelian Barang")
			return
		}
		DetailTRX.IdTransaksi = DataTRX.ID
		DetailTRX.IdBarang = BarangTerjual.ID
		DetailTRX.Qty = uint(Keranjang.JumlahBeli[i])
		DetailTRX.TotalHarga = BarangTerjual.Harga * uint(Keranjang.JumlahBeli[i])

		DetailTRX, err = tc.model.UpdateDetailTransaksi(DetailTRX)
		if err != nil {
			fmt.Println("terjadi masalah ketika update detail data Pembelian")
			return
		}

		DataTRX.TotalQty += DetailTRX.Qty
		DataTRX.TotalHarga += DetailTRX.TotalHarga
	}

	DataTRX, err = tc.model.UpdateTransaksi(DataTRX)
	if err != nil {
		fmt.Println("terjadi masalah ketika update data Pembelian")
		return
	}
	fmt.Print("\033[H\033[2J") //cls

	fmt.Println("BARANG BERHASIL DIBELI :")

	//------------------------------------------- Nota Transaksi START --------------------------

	fmt.Println("\n--------------")
	fmt.Println("NOTA TRANSAKSI")
	fmt.Println("--------------")
	fmt.Println("----------------------------------------------------")
	fmt.Println("ID TANSAKSI :", DataTRX.ID)
	fmt.Println("Nama Petugas :", DataUser.Nama)
	fmt.Println("Nama Customer :", DataCustomer.Nama)
	fmt.Println("----------------------------------------------------")
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	fmt.Fprintln(w, "Nama Barang\tJumlah\tTotal Harga")
	for i, BarangKeranjang := range Keranjang.BarangDibeli {
		//fmt.Fprintln(w, "----------------------------------------------------")
		fmt.Fprintln(w, "-", BarangKeranjang.NamaBarang, "\t", Keranjang.JumlahBeli[i], "\t", uint(Keranjang.JumlahBeli[i])*BarangKeranjang.Harga)
	}
	w.Flush()
	fmt.Println("\n----------------------------------------------------")
	fmt.Println("Total Semua barang :", Keranjang.TotalHarga)
	fmt.Println("----------------------------------------------------")
	//------------------------------------------- Nota Transaksi END --------------------------
	fmt.Println("\nPilih Menu :")
	fmt.Println("[99] KEMBALI KE MENU UTAMA")
	var nextMenu int
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
		case 99:
			fmt.Print("\033[H\033[2J") //cls
			return
		default:
			fmt.Println("- Masukkan input yang valid")
		}

	}

}
