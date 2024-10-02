package main

import "fmt"	


func main() {
var(

	rupiah int
	konversiDollar int
)

fmt.Println("Masukkan rupiah : ")
fmt.Scan(&rupiah)
konversiDollar = rupiah / 15000 //untuk mengonversi rupiah ke dollar 
fmt.Println("Hasil Konversi Dollar adalah : ", konversiDollar)
}