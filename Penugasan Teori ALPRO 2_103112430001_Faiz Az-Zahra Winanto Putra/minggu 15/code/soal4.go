package main

import "fmt"

func main() {

	var x, y, z int
	var belanja, diskon float32

	fmt.Scan(&x, &y, &z)
	diskon = float32(z) * float32(x) / 100

	if diskon > float32(y) {
		diskon = float32(y)
	}

	belanja = float32(z) - diskon
	fmt.Printf("%.0f Rupiah\n", belanja)

}