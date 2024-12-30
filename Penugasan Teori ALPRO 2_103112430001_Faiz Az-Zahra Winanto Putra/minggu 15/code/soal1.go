package main

import "fmt"

func main() {

	var x1, x2 int
	var status bool

	fmt.Scan(&x1, &x2)
	status = (x1%2 == 0 && x2%2 != 0) || (x1%2 != 0 && x2%2 == 0)

	if status == true {
		fmt.Print("Berhasil")
	}

}