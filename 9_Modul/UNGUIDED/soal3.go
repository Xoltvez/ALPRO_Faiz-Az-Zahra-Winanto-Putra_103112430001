package main

import "fmt"

func main() {

	// Membuat variabel bilangan
	var bilangan int

	// Inputan untuk memasukan bilangan
	fmt.Print("Masukan 4 digit bilangan: ")
	fmt.Scan(&bilangan)

	// Rumus untuk menentukan bilagan terurut
	digit1 := bilangan / 1000
	digit2 := (bilangan / 100) % 10
	digit3 := (bilangan / 10) % 10
	digit4 := bilangan % 10

	// Percabangan untuk menentukan apakah bilangan 4 digit atau tidak
	if bilangan >= 1000 && bilangan <= 9999 {
		// Di dalam if terdapat percabangan lagi untuk menentukan bilangan terurut
		if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
			fmt.Print("Bilangan terurut mengecil")
		} else if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
			fmt.Print("Bilangan terurut membesar")
		} else {
			fmt.Print("Bilangan Tidak Terurut")
		}
		// Jika kondisi 4 bilangan tidak terpenuhi, maka akan ditampilkan output harus memasukan 4 digit bilangan.
	} else {
		fmt.Print("Bilangan bukan 4 digit, masukan 4 digit bilangan!")
	}
}