package main

import "fmt"

func main() {
	
	//variable yang terdiri dari luas,keliling,dan sisi
	var (
		luas     int
		keliling int
		sisi     = 27
	)

	//kode untuk menghitung luas dan keliling alun-alun
	luas = sisi * sisi
	keliling = 4 * sisi

	//kode untuk menampilkan output/hasil keliling dan luas yang sudah dihitung
	fmt.Println("Jadi Luas Alun-Alunnya Adalah", luas, "m²")
	fmt.Println("Jadi Keliling Alun-Alunnya Adalah", keliling, "m")
}
