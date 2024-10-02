package main

import "fmt"

func main() {
	//menggunakan 2 variable berupa farenhit dan kelvin
	var (
		farenhit,
		kelvin float32
	)

	//meminta inputan dari user berupa suhu awal farenhit
	fmt.Println("Masukan suhu dalam farenhit")
	fmt.Scan(&farenhit)

	//operasi untuk konversi farenhit ke kelvin
	kelvin = (farenhit-32)*5/9 + 273.15

	//untuk menampilkan output yang sudah dikonversikan ke kelvin
	fmt.Println("Suhu farenhit dalam kelvin adalah", kelvin)
}
