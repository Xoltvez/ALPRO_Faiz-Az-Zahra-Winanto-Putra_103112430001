package main

import "fmt"

func main() {

	var (
		bilanganN, hasil int
	)

	fmt.Print("Masukkan Angka N : ")
	fmt.Scan(&bilanganN)

	for i := 1; i <= bilanganN; i++ {
		hasil = i * i
		fmt.Print(hasil,  " ")

	}

}
