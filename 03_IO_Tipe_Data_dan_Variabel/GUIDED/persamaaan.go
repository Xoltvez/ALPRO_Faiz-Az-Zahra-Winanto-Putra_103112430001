package main

import "fmt"

func main() {

	var x float64
	var persamaan int

	fmt.Println("Masukan Angka: ")

	fmt.Scan(&x)

	// Menghitung persamaan dan mengkonversi hasil ke int
	persamaan = int((2/(x+5)) + 5)

	fmt.Println("Persamaan = ", persamaan)

}
