package main

import "fmt"

func main() {
	var (
		satu, dua, tiga string
		temp            string
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga)
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga + " " + temp)
	temp = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga + " " + temp)
}

//Program ini membaca tiga inputan angka dari pengguna, dengan menampilkan urutan angka tetap sebelum diubah/penukaran
//kemudian mengubah/menukar posisi angka tersebut, dan akhirnya menampilkan urutan angka setelah penukaran sesuai perintah yang ditulis.
