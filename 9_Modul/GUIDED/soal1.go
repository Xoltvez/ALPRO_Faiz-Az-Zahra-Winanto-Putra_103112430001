package main


import "fmt"

func main 	(){
	var umur int
	var ktp bool

	fmt.Print("Masukkan Umur : ")
	fmt.Scan(&umur)

	fmt.Print("Masukan status ktp : ")
	fmt.Scan(&ktp)

	if umur >= 17 && ktp{
		fmt.Print("bisa membuat ktp")
	}else{
		fmt.Print("belum bisa membuat ktp")
	}

}



