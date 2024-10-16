package main

import (
	"fmt"
	"math"
)

func main() {

	// Membuah sebuah variabel yang mencakup volume, jejari, tinggi, angkaN, untuk menyimpan inputan dari user.
	var (
		jejari    float32 //tipe data float32
		tinggi    float32 //tipe data float32
		angkaN int     //tipe data integer
		volume    float32 //tipe data float32
	)
	//meminta user untuk memasukan angkaN
	fmt.Print("Masukan angka N : ")
	fmt.Scan(&angkaN)

	//perulangan untuk menghitung volume kerucut sesuai angka yang diinputkan oleh user
	for i := 1; i <= angkaN; i++ {

		//meminta user untuk memasukan jejari
		fmt.Println("Masukan jejari : ")
		fmt.Scan(&jejari)

		//meminta user untuk memasukan tinggi
		fmt.Println("Masukan tinggi: ")
		fmt.Scan(&tinggi)

		//menghitung volume kerucut
		volume = (1.0 / 3.0) * math.Pi * jejari * jejari * tinggi
		//menampilkan hasil perhitungan volume kerucut
		fmt.Println("Hasil volume kerucut adalah  ", volume)
	}

}
