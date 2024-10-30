package main

import "fmt"

func main() {
	// Variabel untuk menyimpan data
	var angka int
	var hasil string

	// Inputan dari user untuk memasukan angka
	fmt.Print("Masukan angka : ")
	fmt.Scan(&angka)

	//percabangan jika angka < 0 dan angka habis dibagi 2 maka menampilkan "Genap Negatif"
	//percabangan jika angka < 0 dan angka tidak habis dibagi 2 maka menampilkan "Bukan"
	if angka < 0 && angka%2 == 0 {

		hasil = "Genap Negatif"

	} else {
		hasil = "Bukan"
	}

	//menampilkan output
	fmt.Print(hasil)

}
