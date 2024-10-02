package main

import "fmt"

func main() {
	var (
		farenhit,
		reamur,
		kelvin,
		celcius float32
	)

	fmt.Println("Masukan Suhu Celcius")

	fmt.Scan(&celcius)

	//kode untuk mengkonversi suhu 
	reamur = celcius * 4 / 5
	farenhit = (celcius * 9/5) + 32
	kelvin = celcius + 273


	//kode untuk menampilkan hasil konversi dari celcius ke reamur, farenhit, dan kelvin
	fmt.Println("Derajat Reamur :", reamur)
	fmt.Println("Derajat Farenhit :", farenhit)
	fmt.Println("Derajat Kelvin :", kelvin)
}
