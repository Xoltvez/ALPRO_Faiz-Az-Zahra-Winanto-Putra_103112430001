package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data mencakup a, pangkat, dan hasil
	var a, pangkat, hasil int // dengan tipe data integer

	// meminta user untuk memasukan angka dan pangkat
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a)

	fmt.Print("Masukan pangkat: ")
	fmt.Scan(&pangkat)

	// Untuk mendefinisikan nilai dari hasil
	hasil = 1

	// Perulangan untuk membuat operasi perpangkatan
	for i := 1; i <= pangkat; i++ {
		hasil *= a
	}

	// Menampilkan hasil dari operasi perpangkatan
	fmt.Print(hasil)
}