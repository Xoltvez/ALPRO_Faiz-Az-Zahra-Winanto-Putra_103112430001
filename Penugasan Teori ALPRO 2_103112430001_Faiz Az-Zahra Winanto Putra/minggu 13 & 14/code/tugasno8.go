package main

import "fmt"

func main() {

	var angka int

	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {

		for j := 1; j <= angka; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()

	}

}