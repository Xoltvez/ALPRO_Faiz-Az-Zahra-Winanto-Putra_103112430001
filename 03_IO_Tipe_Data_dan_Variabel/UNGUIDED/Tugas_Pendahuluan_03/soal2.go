package main

import (
	"fmt"
)

func main() {
	//variable yang terdiri dari angka kedua, angka pertama, dan operator
	var angkaKedua, angkaPertama float64
	var operator string

	// Menampilkan menu kalkulator, kemudian meminta input angka pertama dari user
	fmt.Println("Kalkulator Sederhana")
	fmt.Println("Masukkan angka yang ingin dioperasikan:")
	fmt.Print("Masukan Angka pertama: ")
	fmt.Scanln(&angkaPertama)

	// Meminta input operator dari user
	fmt.Println("Pilih operasi: +, -, *, /")
	fmt.Print("Operator: ")
	fmt.Scanln(&operator)
	
	//Meminta input angka kedua dari user
	fmt.Print("Masukan Angka kedua: ")
	fmt.Scanln(&angkaKedua)

	

	// Melakukan operasi berdasarkan input pengguna
	switch operator {
	case "+":
		fmt.Printf("%.2f + %.2f = %.2f\n", 	angkaPertama, angkaKedua, angkaPertama+angkaKedua)
	case "-":
		fmt.Printf("%.2f - %.2f = %.2f\n", angkaPertama, angkaKedua, angkaPertama-angkaKedua)
	case "*":
		fmt.Printf("testing", angkaPertama, angkaKedua, angkaPertama*angkaKedua)
	case "/":
		if angkaKedua != 0 {
			fmt.Printf("%.2f / %.2f = %.2f\n", angkaPertama, angkaKedua, angkaPertama/angkaKedua)
		} else {
			fmt.Println("Error: Pembagian dengan nol tidak diperbolehkan.")
		}
	default:
		fmt.Println("Error: Operator tidak valid.")
	}
}