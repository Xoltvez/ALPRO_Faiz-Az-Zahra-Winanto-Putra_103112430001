package main

import "fmt"

func main() {

	var bilangan int

	fmt.Print("Masukan 4 digit bilangan: ")
	fmt.Scan(&bilangan)

	digit1 := bilangan / 1000
	digit2 := (bilangan / 100) % 10
	digit3 := (bilangan / 10) % 10
	digit4 := bilangan % 10

	if bilangan >= 1000 && bilangan <= 9999 {
		if digit1 > digit2 && digit2 > digit3 && digit3 > digit4 {
			fmt.Print("Bilangan terurut mengecil")
		} else if digit1 < digit2 && digit2 < digit3 && digit3 < digit4 {
			fmt.Print("Bilangan terurut membesar")
		} else {
			fmt.Print("Bilangan Tidak Terurut")
		}
	} else {
		fmt.Print("Bilangan bukan 4 digit, masukan 4 digit bilangan!")
	}
}