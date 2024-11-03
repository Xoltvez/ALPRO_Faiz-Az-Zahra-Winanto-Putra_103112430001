package main

import "fmt"


func main(){

	//variable untuk menyimpan data dengan tipe data string
	var huruf string
	var hasil string

	//meminta inputan huruf vokal dari user
	fmt.Print("Masukkan Huruf Vokal : ")
	fmt.Scan(&huruf)

	//percabangan jika user memasukkan huruf= "A,I,U,E,O" dan "a,i,u,e,o" merupakan huruf vokal
	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" || huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		hasil = "Huruf Vokal"

		//jika user memasukan tidak selain = "A,I,U,E,O" dan "a,i,u,e,o" merupakan huruf konsonan
	} else {
		hasil = "Huruf Konsonan"
	}
	//menampilkan output
	fmt.Println(hasil)
}