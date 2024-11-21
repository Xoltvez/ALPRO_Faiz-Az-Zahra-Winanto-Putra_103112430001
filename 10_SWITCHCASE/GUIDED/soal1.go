package main

import (
	"fmt"
)


func main () {
	//deklarasikan variabel jam12 dan jam24 tipe data integer, dan keterangan tipe data string
	var jam12, jam24 int
	var keterangan string
	//meminta input user untuk memasukan jam
	fmt.Scan(&jam24)

	//menggunakan switch case untuk menentukan jam dari kondisi tertentu
	switch {
	case jam24 == 0:
		jam12 = 12
		keterangan = "AM"

	case jam24 < 12:
		jam12 = jam24
		keterangan = "AM"

	case jam24 == 12:
		jam12 = 12
		keterangan = "PM"

	case jam24 > 12 && jam24 < 12:
		jam12 = jam24 - 12
		keterangan = "PM"
		

		//menampilkan kondisi yg semua tidak terpenuhi
		default:
			fmt.Println("Jam tidak valid")
	}
	//menampilkan hasil
	fmt.Println(jam12, keterangan)
}