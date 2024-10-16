package main

import "fmt"

func main() {

	var (
		// Variabel untuk menyimpan inputan angka dari user
		angkaN int
	)

	//meminta inputan dari user untuk memasukan angka
	fmt.Print("Masukan angka yang ingin dimasukan: ")
	fmt.Scan(&angkaN)

	//perulangan untuk menghitung penjumlahan dari deretan angka yang diinputkan
	hasil := 0
	for i := 1; i <= angkaN; i++ {
		hasil += i
	}

	//menampilkan hasil penjumlahan deret
	fmt.Print("Hasil jumlah dari deret adalah  ", hasil)
}
