package main

import "fmt"

func main() {

	// Variabel untuk menyimpan data menggunakan tipe data integer
	var (
		jumlah int
		total int
	)

	// Meminta user untuk memasukan jumlah barang
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlah)

	// Perulangan untuk menentukan total poin (10 poin untuk 1-5, 15 poin untuk 6-n)
	for i := 1; i <= jumlah; i++ {
		if i <= 5 {
			total += 10
		} else {
			total += 15
		}
	}

	// Menampilkan total poin yang didapatkan
	fmt.Print("Total poin yang didapat: ", total)
}