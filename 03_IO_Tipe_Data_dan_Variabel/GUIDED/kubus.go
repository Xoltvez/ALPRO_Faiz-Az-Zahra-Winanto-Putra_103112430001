package main

import (
	"fmt"
)

func main() {

	var (
		sisi,
		volume int
	)
	fmt.Println("Masukkan sisi kubus : ")
	fmt.Scan(&sisi)
	volume = sisi * sisi * sisi
	fmt.Println("Volume kubus = ", volume)
}
