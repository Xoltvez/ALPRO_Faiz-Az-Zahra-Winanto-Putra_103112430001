package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	// Input tiga bilangan
	fmt.Scan(&a)
	fmt.Scan(&b) 
	fmt.Scan(&c)

	// Cari nilai terbesar dan terkecil
	max := a
	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	min := a
	if b < min {
		min = b
	}
	if c < min {
		min = c
	}

	// Tampilkan hasil
	fmt.Printf("Bilangan terbesar: %d\n", max)
	fmt.Printf("Bilangan terkecil: %d\n", min)
}
