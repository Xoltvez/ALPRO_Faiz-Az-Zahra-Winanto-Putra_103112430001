package main

import "fmt"


func main (){

	var (
		angkaPertama, angkaKedua, hasil int
	)


	fmt.Print("Masukan Angka Pertama : ")
	fmt.Scan(&angkaPertama)

	fmt.Print("masukan angka kedua : ")
	fmt.Scan(&angkaKedua)

	for i := 1; i <= angkaPertama; i+=1{
		hasil = hasil + angkaKedua
	}

	fmt.Print("Hasilnya adalah : ", hasil)
}