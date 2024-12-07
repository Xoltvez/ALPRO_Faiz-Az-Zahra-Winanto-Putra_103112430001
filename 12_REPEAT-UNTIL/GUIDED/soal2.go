package main

import "fmt"

func main (){

	var angka int

	for done := false; !done; {

		fmt.Scan(&angka)

		done = (angka > 0)

	}

	fmt.Println(angka, "adalah bilangan bulat positif")


}