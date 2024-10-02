package main

import "fmt"

func main() {
	var r, luas float64
	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&r)
	luas = 3.14 * r * r
	fmt.Print("Luas lingkaran = ", luas)
}
