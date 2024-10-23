package main

import "fmt"

func main(){

	var (
		dinar, dirham, fals, qirat int
	)

	fmt.Print("Masukkan qirat: ")
	fmt.Scan(&qirat)

	fals = qirat / 6
    qirat = qirat % 6

    dirham = fals / 10
    fals = fals % 10

    dinar = dirham / 10
    dirham = dirham % 10


	fmt.Println("Hasil konversi:")
    fmt.Printf("Dinar: %d\n", dinar)
    fmt.Printf("Dirham: %d\n", dirham)
    fmt.Printf("Fals: %d\n", fals)
    fmt.Printf("Qirat: %d\n", qirat)
}