package main

import "fmt"

func main() {
	//variable untuk menyimpan jari-jari, luas, dan keliling
	var r float64

	//meminta input dari user untuk memasukan jari-jari
	fmt.Print("Masukan jari-jari lingkaran: ")
	fmt.Scan(&r)

	//rumus luas dan keliling
	var luas float64 = 3.14 * r * r
	var keliling float64 = 2 * 3.14 * r

	//menampilkan output
	fmt.Println("Hasil dari luas lingkaran dengan jari-jari:", r, "adalah", luas, "dan kelilingnya", keliling)
}
