package main

import (
	"fmt"
)

func main(){
	var (
		hargaAwal, diskon, hargaAkhir int
	)

	fmt.Println("Masukan Harga Awal :")
	fmt.Scan(&hargaAwal)

	fmt.Println("Masukan Diskon :")
	fmt.Scan(&diskon)

	hargaAkhir = (hargaAwal * diskon) /100
	hargaAkhir = hargaAwal - hargaAkhir
	fmt.Println("Harga setelah di diskon adalah", hargaAkhir)
}