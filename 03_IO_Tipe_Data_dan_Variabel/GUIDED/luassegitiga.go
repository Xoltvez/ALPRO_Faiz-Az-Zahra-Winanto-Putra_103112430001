package main

import (
	"fmt"
)

func main() {


	var(
		alas,
		tinggi int
		luasSegitiga int
	)


	fmt.Println("Masukkan alas segitiga : ")
	fmt.Scan(&alas)
	fmt.Println("Masukkan tinggi segitiga : ")
	fmt.Scan(&tinggi)
	luasSegitiga = alas * tinggi / 2
	fmt.Println("Luas segitiga = ", luasSegitiga)

}