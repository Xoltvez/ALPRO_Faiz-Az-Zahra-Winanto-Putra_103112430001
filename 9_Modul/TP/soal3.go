package main

import "fmt"

func main() {

	//variable untuk menyimpan data dengan tipe data integer dan string
	var umur int
	var kewarganegaraan string

	//meminta inputan dari user untuk memasukkan umur dan kewarganegaraan
	fmt.Print("Masukkan Umur :  ")
	fmt.Scan(&umur)

	fmt.Print("Masukkan Kewarganegaraan :  ")
	fmt.Scan(&kewarganegaraan)

	//percabangan jika umur memenuhi >= 17 dan kewarganegaraan = "indonesia" akan menampilkan "anda bisa mengikuti pemilu"
	if umur >= 17 && kewarganegaraan == "indonesia" {
		fmt.Println("Anda bisa mengikuti pemilu.")

		//lain jika tidak memenuhi syarat umur dan kewarganegaraan = "indonesia" akan menampilkan jika umur < 17 dan kewarganegaraan = "indonesia" akan menampilkan "anda tidak bisa mengikuti pemilu karena belum cukup umur"
	} else {
		if umur < 17 {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena belum cukup umur.")
		}
		//jika kewarganegaraan bukan = "indonesia" akan menampilkan "anda tidak bisa mengikuti pemilu karena bukan warga negara Indonesia"
		if kewarganegaraan != "indonesia" {
			fmt.Println("Anda tidak bisa mengikuti pemilu karena bukan warga negara Indonesia.")
        }
	}
}
