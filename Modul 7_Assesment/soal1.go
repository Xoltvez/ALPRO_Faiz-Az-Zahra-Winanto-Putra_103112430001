package main

import"fmt"


func main(){

	var (
		angkaX, angkaY int
	)

	fmt.Print("Masukkan Bilangan Bulat Positif Y : ")
	fmt.Scan(&angkaX)

	fmt.Print("Masukkan Bilangan Bulat Positif X : ")
	fmt.Scan(&angkaY)
	
	hasil := 0
	for i := angkaX; i <= angkaY; i++ {
		hasil += i
	}

	fmt.Print("Hasilnya penjumlahan semua bilangan X dan Y adalah : ", hasil)


}

