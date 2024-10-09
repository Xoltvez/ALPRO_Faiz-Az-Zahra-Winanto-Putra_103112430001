package main

import (
	"fmt"
)

func main() {

	var (
		beratBadan,
		tinggiBadan,
		bmi float32
	)

	fmt.Print("Masukkan berat badan: ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukkan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	tinggiBadan = tinggiBadan / 100

	bmi = beratBadan / (tinggiBadan * tinggiBadan)

	fmt.Printf("Nilai BMI: %.f\n", bmi)
}
