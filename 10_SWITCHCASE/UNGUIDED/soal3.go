package main

import (
	"fmt"
)

func main() {
	//dekreasi variabel bilangan dengan tipe data integer
	var bilangan int

	//input user untuk memasukan bilangan
	fmt.Print("Masukkan sebuah bilangan bulat: ")
	fmt.Scanln(&bilangan)

	// Gunakan switch case untuk memeriksa kategori
	switch {
	case bilangan%2 != 0:
		// // untuk bilangan dengan kategori bilangan ganjil
		result := bilangan + (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Ganjil\n")
		fmt.Printf("Hasil penjumlahan dengan bilangan berikutnya %d + %d = %d\n", bilangan, bilangan+1, result)

	case bilangan%2 == 0:
		// untuk bilangan dengan kategori bilangan genap
		result := bilangan * (bilangan + 1)
		fmt.Printf("Kategori: Bilangan Genap\n")
		fmt.Printf("Hasil perkalian dengan bilangan berikutnya %d * %d = %d\n", bilangan, bilangan+1, result)
	}

	if bilangan%5 == 0 {
		// untuk bilangan dengan kategori bilangan kelipatan 5
		result := bilangan * bilangan
		fmt.Printf("Kategori: Bilangan Kelipatan 5\n")
		fmt.Printf("Hasil kuadrat dari %d ^ 2 = %d\n", bilangan, result)
	}

	if bilangan%10 == 0 {
		// // untuk bilangan dengan kategori bilangan kelipatan 10
		result := bilangan / 10
		fmt.Printf("Kategori: Bilangan Kelipatan 10\n")
		fmt.Printf("Hasil pembagian antara %d / 10 = %d\n", bilangan, result)
	}
}
