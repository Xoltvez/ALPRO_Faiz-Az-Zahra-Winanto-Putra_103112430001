package main

import "fmt"

func main() {

	//variable untuk menyimpan data dengan tipe data string dan integer
	var nilai int
	var indeks string

	//meminta inputan nilai dari user
	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&nilai)

	//percabangan jika nilai lebih dari 90 indeks = "A"
	if nilai > 90 {
		indeks = "A"

		//jika nilai >= 80 indeks = "AB"
	} else if nilai >= 80 {
		indeks = "AB"

		//jika nilai >= 70 indeks = "B"
	} else if nilai >= 70 {
		indeks = "B"

		//jika nilai <70 indeks = "C"
	} else if nilai < 70 {
		indeks = "C"
	}

	//menampilkan output
	fmt.Println(indeks)
}
