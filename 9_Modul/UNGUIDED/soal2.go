// 1. Program yang dberikan belum sesuai, sehingga program belum bisa berjalan dan tidak dapat menginputkan nilai yang telah diberikan.

// 2. Kesalahan dalam program diatas adalah
// a. Penulisan nama variable yang belum sesuai, seperti pada percabangan yang menyimpan nilai NMK, seharunya menggunakan variabel NMK
// b. Sebenarnya masih benar-benar saja, namun kurang efektif yaitu pada variabel NAM menggunakan float64, karena output dalam nilai mahasiswa lebih efektif menggunakan float32
// c. Dalam penggunaan percabangan, penulisan kondisi masih belum tepat.
// d. Penggunaan tanda kutip yang kurang tepat ("")

// 3. Berikut merupakan pembenaran codingannya.

package main

import "fmt"

func main() {

	// Membuat variable
	var nam float32
	var nmk string

	// Membuat sebuah inputan untuk memasukan nilai mata kuliah
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Percabangan untuk menentukan nilai masuk ke kategori mana
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam <= 50 && nam > 40 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	// Menampilkan kategori nilai
	fmt.Println("Nilai akhir mata kuliah: ", nmk)

}