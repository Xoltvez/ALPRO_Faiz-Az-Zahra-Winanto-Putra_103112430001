package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data mencakup angka(a), dan hasil
	var a, hasil int // tipe data integer

	// Inputan untuk memasukan angka dan pangkat
	fmt.Print("Masukan angka: ")
	fmt.Scan(&a)

	// Untuk mendefinisikan nilai dari hasil
	hasil = 1

	// Perulangan untuk membuat operasi perpangkatan
	for i := 1; i <= a; i++ {
		hasil *= i
	}

	// Menampilkan hasil dari operasi perpangkatan
	fmt.Print(hasil)
}