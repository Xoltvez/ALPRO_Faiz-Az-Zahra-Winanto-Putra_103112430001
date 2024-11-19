package main

import "fmt"


func main(){

	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	if huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" || huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" {
		fmt.Print("Huruf Vokal")
	}else{
		if huruf >= "a" && huruf <= "z" || huruf >= "A" && huruf <="Z"{
		fmt.Print("konsonan")
		}else{
			fmt.Print("bukan huruf")
		}
	}
}