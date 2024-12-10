package main

import "fmt"

func main() {
	//deklarasi variable angka dan prima dengan tipe data integer dan boolean
	var (
		angka int
		prima bool
	)

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan pertama untuk mengurutkan angka dari 2 hingga angka yang dimasukan
	for i := 2; i <= angka; i++ {
		prima = true
		// Perulangan kedua untuk memeriksa apakah angka tersebut prima
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				prima = false
				break
			}
		}
		if prima {
			fmt.Print(i, " ")
		}
	}
}