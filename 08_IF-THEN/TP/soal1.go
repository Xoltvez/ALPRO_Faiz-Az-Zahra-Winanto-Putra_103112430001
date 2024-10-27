package main

import "fmt"

func main() {

	//variable untuk menyimpan data dengan var nilaiUjian tipe data integer dan keterangan tipe data string
	var nilaiUjian int
	var keterangan string
	
	//meminta inputan dari user
	fmt.Print("Masukkan Nilai Ujian : ")
	fmt.Scan(&nilaiUjian)

	//percabangan jika nilai lebih dari 70 maka "keterangan" lulus
	if nilaiUjian >= 70 {
		keterangan = "Lulus"
	//kalau nilai ujian dibawah 70 maka "keterangan" tidak lulus
	} else {
		keterangan = "Tidak Lulus"
	}

	//menampilkan output
	fmt.Print(keterangan)

}
