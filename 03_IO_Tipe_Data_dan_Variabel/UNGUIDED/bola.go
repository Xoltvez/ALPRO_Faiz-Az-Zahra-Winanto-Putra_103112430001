package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		jari,
		volumeBola,
		luasBola float64
	)

	fmt.Println("Masukkan jari-jari bola: ")
	fmt.Scan(&jari)

	//untuk menghitung volume bola dan luas bola
	volumeBola = (4.0 / 3.0) * math.Pi * math.Pow(jari, 3)
	luasBola = 4 * math.Pi * math.Pow(jari, 2)

	//untuk menampilkan hasil operasi yang sudah dihitung
	fmt.Printf("Volume bola = %.4f\n", volumeBola)
	fmt.Printf("Luas bola = %.4f\n", luasBola)
}

//math.pow adalah untuk menghitung pangkat dan math.pi adalah untuk menghitung nilai pi
