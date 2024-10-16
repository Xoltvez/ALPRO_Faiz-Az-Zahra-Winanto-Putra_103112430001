package main

import "fmt"

func main (){

	var(
		a, b int

	)


	fmt.Print("Masukkan angka pertama :")
	fmt.Scan(&a)

	fmt.Print("Masukkan angka kedua :")	
	fmt.Scan(&b)

	for i:= a; i <= b; i++ {
		fmt.Print(i, " ")
	}
}