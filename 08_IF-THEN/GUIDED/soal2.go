package main

import "fmt"

func main() {

	//variabel untuk menyimpan data dengan var angka tipe data integer dan keterangan tipe data string
	var angka int
	var keterangan string

	//meminta inputan dari user
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&angka)


	if angka < 1 {
		keterangan = "Bukan Positif"

	
	} else {
		keterangan = "Positif"
	}

	//menampilkan output
	fmt.Print(keterangan)

}
