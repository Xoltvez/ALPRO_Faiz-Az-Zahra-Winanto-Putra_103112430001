package main

import "fmt"

func main() {

	// Variabel untuk menyimpan range yang merupakan bilangan genap
	var (
		bilanganGenap int
	)

	// Meminta user untuk memasukan range angka
	fmt.Print("Masukkan Angka: ")
	fmt.Scan(&bilanganGenap)

	// Menampilkan bilangan genap dari angka yang dimasukan menggunakan perulangan yang analogikan angka yang bisa dibagi dengan 2.
	for i := 1; i <= bilanganGenap; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
}