package main

import "fmt"

func main() {

	// Membuat variable
	var bilangan, b, result int

	// Inputan untuk memasukan angka
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&bilangan)

	// Switch case untuk menentukan kategori dari angka yang dimamsukan
	switch {
	// Digunakan untuk menentukan apakah kelipatan 10
	case bilangan%10 == 0:
		result = bilangan / 10

		fmt.Println("Kategori: Kelipatan 10")
		fmt.Println("result pembagian dengan bilangan berikutnya ", bilangan, "/", "10 =", result)
	// Menentukan apakah kelipatan 5
	case bilangan%5 == 0 && bilangan > 5:
		result = bilangan * bilangan

		fmt.Println("Kategori: Kelipatan 5")
		fmt.Println("result perkalian dengan bilangan berikutnya ", bilangan, "^2", " = ", result)
	// Menentukan apakah genap
	case bilangan%2 == 0:
		b = bilangan + 1
		result = bilangan * b

		fmt.Println("Kategori: Genap")
		fmt.Println("result perkalian dengan bilangan berikutnya ", bilangan, "*", b, " = ", result)
	// Menentukan apakah ganjil
	default:
		b = bilangan + 1
		result = bilangan + b

		fmt.Println("Kategori: Ganjil")
		fmt.Println("result penjumlahan dengan bilangan berikutnya ", bilangan, "+", b, " = ", result)
	}

}