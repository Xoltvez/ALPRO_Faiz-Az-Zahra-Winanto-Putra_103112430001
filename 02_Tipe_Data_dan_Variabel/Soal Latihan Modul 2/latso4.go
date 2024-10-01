package main

import "fmt"

func main() {
	var(
		
		konversiCelcius int
		farenhit int
	)

	fmt.Println("Masukkan suhu dalam fahrenheit:")
	fmt.Scan(&farenhit)
	
	konversiCelcius = (farenhit - 32) * 5 / 9 //untuk konversi celcius
	fmt.Println("Suhu dalam celcius adalah:", konversiCelcius, "°C")
}

