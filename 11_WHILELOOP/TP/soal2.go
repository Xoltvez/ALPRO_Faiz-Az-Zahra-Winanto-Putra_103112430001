package main

import "fmt"

func main() {

	// Membuat variabel
	var hargaBarang, totalHarga int

	// Perulangan
	for {

		// Inputan untuk memasukan harga barang
		fmt.Print("Silahkan Masukan Harga Belanja Anda (Selesai masukan 0): ")
		fmt.Scan(&hargaBarang)

		// Disini terdapat kondisi percabangan, yaitu jika user menginputkan nilai 0, maka perulangan akan berhenti.
		if hargaBarang == 0 {
			break
		}

		// Total harga akan dijumlah semua ketika sudah selesai.
		totalHarga += hargaBarang
	}

	// Menampilkan total harga belanjaan
	fmt.Print("Total belanja adalah:", totalHarga)
}