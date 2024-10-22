package main

import "fmt"

func main() {
	// Variabel untuk menyimpan data
	var N int

	// Inputan untuk memasukan nilai n
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&N)


	// Mengalikan angka perulangan dengan angka itu sendiri
	for i := 1; i <= N; i++ {
		fmt.Println(i * i)
	}
}