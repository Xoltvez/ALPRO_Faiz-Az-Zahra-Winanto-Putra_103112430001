package main

import (
	"fmt"
)

func main() {

	// Membuah sebuah variabel bernama angkaN untuk menyimpan inputan dari user.
	var (
		bilanganN, hasil int
	)

	// Meminta user untuk menginputan bilangan bulat.
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&bilanganN)

	// Rumus untuk mencari jumlah deret angka menggunakan perulangan
	hasil = 0
	for i := 1; i <= bilanganN; i++ {
		hasil += i
	}

	// Output atau hasil dari perulangan diatas
	fmt.Print("Hasil jumlah dari deret 1 hingga ", bilanganN, " adalah ", hasil)
}