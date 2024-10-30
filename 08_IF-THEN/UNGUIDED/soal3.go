package main

import (
	"fmt"
)

func main() {

	// Membuat variabel
	var x, y int
	var hasilX, hasilY bool

	// Inputan untuk memasukan nilai x dan y
	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)

	// Terdapat dua percabangan, yaitu untuk mengetahui apakah x adalah faktor dari y, dan apakah y adalah faktor dari x.
	if y%x == 0 {
		hasilX = true
	}

	if x%y == 0 {
		hasilY = true
	}

	// Menampilkan hasil yang terpenuhi
	fmt.Println(hasilX)
	fmt.Println(hasilY)
}