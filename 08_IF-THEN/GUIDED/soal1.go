package main

import "fmt"

func main (){

	var angka int
	

	fmt.Print("Masukan angka : ")
	fmt.Scan(&angka)


		
	if angka < 0{
		angka = angka * -1
	}
	fmt.Print("Nilai mutlaknya adlaah  ", angka )
}