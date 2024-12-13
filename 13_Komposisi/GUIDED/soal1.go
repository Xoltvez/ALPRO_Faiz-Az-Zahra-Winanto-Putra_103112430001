package main

import "fmt"

func main(){

	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan untuk memeriksa faktor dari angka yang dimasukan
	for i := 1; i <= angka; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}