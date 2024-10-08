package main

import (
	"fmt"
)

func hitungGaji(jamKerja float64, upahPerJam float64) float64 {
	const batasJamNormal = 40
	const faktorLembur = 1.5
	const mingguDalamBulan = 4

	// Menggunakan operator ternary-like dengan if-else
	jamNormal := jamKerja
	jamLembur := 0.0

	if jamKerja > batasJamNormal {
		jamNormal = batasJamNormal
		jamLembur = jamKerja - batasJamNormal
	}

	// Menghitung total gaji mingguan
	gajiMingguan := (jamNormal * upahPerJam) + (jamLembur * faktorLembur * upahPerJam)

	// Menghitung total gaji bulanan
	return gajiMingguan * mingguDalamBulan
}

func main() {
	// Menggunakan tipe data yang lebih general
	var jamKerjaMingguan, upahPerJam float64

	// Meminta input dari pengguna dengan format string yang lebih deskriptif
	fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
	fmt.Scanln(&jamKerjaMingguan)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&upahPerJam)

	// Menghitung total gaji bulanan
	totalGajiBulanan := hitungGaji(jamKerjaMingguan, upahPerJam)

	// Menampilkan hasil gaji bulanan dengan format output yang lebih jelas
	fmt.Printf("Total gaji bulanan yang diperoleh adalah: Rp%.2f\n", totalGajiBulanan)
}
