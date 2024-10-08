package main

import (
	"fmt"
)

func main() {
	var (
		//variable untuk menyimpan jari-jari, luas, keliling, dan phi
		luas     float64
		jejari   float64
		phi      = 3.14
		keliling float64
	)

	fmt.Println("Masukkan jari-jari:")

	//meminta input dari user untuk jari-jari
	fmt.Scan(&jejari)

	luas = phi * (jejari * jejari)               //rumus luas lingkaran
	keliling = 2 * phi * jejari                  // rumus keliling lingkaran
	fmt.Println("Luas lingkaran:", luas)         //menampilkan hasil luas yang sudah dihitung
	fmt.Println("Keliling lingkaran:", keliling) //menampilkan hasil keliling yang sudah dihitung

}
