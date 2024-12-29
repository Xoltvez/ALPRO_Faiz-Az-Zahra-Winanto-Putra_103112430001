package main

import "fmt"

func main() {

	var angka, baris, kolom int

	fmt.Scan(&angka)

	for baris = 1; baris <= angka; baris++ {

		for kolom = 1; kolom <= angka; kolom++ {
			if baris == kolom || baris+kolom == angka+1 {
				fmt.Print(baris, " ")
			} else {
				fmt.Print("  ")
			}
		}

		fmt.Println()
	}
}