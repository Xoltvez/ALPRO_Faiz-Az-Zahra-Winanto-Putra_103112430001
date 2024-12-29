package main

import "fmt"

func main() {

	
	var huruf string

	
	fmt.Print("Masukan huruf: ")
	fmt.Scan(&huruf)

	
	if huruf == "a" || huruf == "i" || huruf == "u" || huruf == "e" || huruf == "o" || huruf == "A" || huruf == "I" || huruf == "U" || huruf == "E" || huruf == "O" {
		fmt.Println("Huruf Vokal")
		
	} else {
		
		if huruf >= "a" && huruf <= "z" || huruf >= "A" && huruf <= "Z" {
			fmt.Println("Huruf Konsonan")
			
		} else {
			fmt.Println("Bukan Huruf")
		}
	}
}