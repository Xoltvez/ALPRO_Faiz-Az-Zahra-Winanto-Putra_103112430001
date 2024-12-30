package main

import "fmt"

func main() {
	var tarif, potongan, diskon, tarifAwal float32
	var durasi, kelebihan int
	var member string

	fmt.Scan(&member, &durasi)

	if member == "Gold" {
		diskon = 0.5
	} else if member == "Silver" {
		diskon = 0.25
	} else {
		diskon = 0
	}

	if durasi <= 2 {
		tarifAwal = float32(durasi) * 65000
	} else {
		kelebihan = durasi - 2
		tarifAwal = (2 * 65000) + (20000 * float32(kelebihan))
	}

	potongan = diskon * tarifAwal
	tarif = tarifAwal - potongan

	if member == "None" && durasi <= 2 {
		fmt.Printf("IDR %.0f\n", tarif)
	} else if member == "Gold" || member == "Silver" {
		fmt.Printf("IDR %.0f\n", tarif)
	} else {
		fmt.Printf("IDR %.0f\n", tarif)
	}
}