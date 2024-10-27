package main

import "fmt"

func main() {

	//variabel untuk menyimpan data dengan var angka tipe data integer dan keterangan tipe data string
	var angka int
	var keterangan string

	//meminta inputan dari user
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&angka)

	//percabangan jika angka bisa habis dibagi dengan 2 maka menmpilkan "Merupakan angka genap"
	if angka%2 == 0 {
		keterangan = "Merupakan angka genap"

	//kalau angka tidak bisa habis dibagi dengan 2 maka menmpilkan "Merupakan angka ganjil"
	} else {
		keterangan = "Merupakan angka ganjil"
	}

	//menampilkan output
	fmt.Print(keterangan)

}
