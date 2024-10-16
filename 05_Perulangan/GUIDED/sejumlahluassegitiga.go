package main

import "fmt"


func main(){


	var(
		count, alas, tinggi int

	)
	fmt.Println("Masukan yang ingin dihitung : ")
	fmt.Scan(&count)


	for i := 1; i <= count; i++ {

		fmt.Print("Masukan alas segitiga: ")
		fmt.Scan(&alas)
		fmt.Print("Masukan tinggi segitiga: ")
		fmt.Scan(&tinggi)
		fmt.Println("Luas segitiga = ", alas * tinggi / 2)
		
	}
}