package main

import "fmt"

func main() {

	//variabel untuk menyimpan data dengan var angka tipe data integer dan keterangan tipe data string
	var angka int
	var keterangan string

	//meminta inputan dari user
	fmt.Print("Masukkan Angka : ")
	fmt.Scan(&angka)

	if angka < 0  && angka%2 == 0 {
		keterangan = "True"

	}else{
		keterangan = "False"
	}
			

	//menampilkan output	
	fmt.Print(keterangan)

}
