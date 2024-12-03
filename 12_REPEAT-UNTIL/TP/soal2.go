package main

import "fmt"

func main() {

	var kata string


	for {
		fmt.Print("Masukkan kata:")
		fmt.Scan(&kata)

		fmt.Println("Anda mengetik:", kata)
		if kata == "telkom" {
			fmt.Println("Program Selesai.")
			break
		} 
	}
}