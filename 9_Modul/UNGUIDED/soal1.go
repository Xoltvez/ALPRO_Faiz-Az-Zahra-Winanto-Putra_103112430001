package main

import "fmt"

func main() {
	var (
		gram, kg, total int
	)

	fmt.Print("Masukan berat (Satuan Gram): ")
	fmt.Scan(&gram)

	kg = gram / 1000
	gram = gram % 1000
	total = kg * 10000

	if kg <= 10 {
		if gram >= 500 {
			total = total + (gram * 5)
		} else {
			total = total + (gram * 15)
		}
		fmt.Print(total)
	} else {
		fmt.Println(total)
	}
}