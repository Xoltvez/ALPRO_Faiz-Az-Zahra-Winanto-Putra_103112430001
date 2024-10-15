package main

import (
	"fmt"
)

func main() {
	// Meminta user memasukkan jumlah baris yang diinginkan
	var jumlahBaris int
	fmt.Print("Masukkan jumlah baris segitiga: ")
	fmt.Scan(&jumlahBaris)

	// Melakukan cek untuk setiap baris seusai jumlah baris yang diinoutkan oleh user
	for baris := 1; baris <= jumlahBaris; baris++ {
		// Di setiap baris, cetak bintang sebanyak nomor barisnya
		cetakBintang(baris)
		// Pindah ke baris baru setelah mencetak bintang
		fmt.Println()
	}
}

// Fungsi untuk mencetak sejumlah bintang pada satu baris
func cetakBintang(n int) {
	for j := 0; j < n; j++ {
		fmt.Print("* ")
	}
}
