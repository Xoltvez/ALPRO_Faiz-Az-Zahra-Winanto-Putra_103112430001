package main

import "fmt"

func main() {

	var hargaBarang, total int

	for{
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&hargaBarang)

		if hargaBarang == 0 {
			break
		}
		total += hargaBarang
	}
	fmt.Println("Total belanja Anda:", total)
}