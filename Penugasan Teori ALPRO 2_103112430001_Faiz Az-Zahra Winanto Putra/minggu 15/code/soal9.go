package main

import "fmt"

func main() {

	var p, r, hari int

	fmt.Scan(&p, &r)
	hari = p / r

	if p%r > 0 {
		hari++
	}

	fmt.Print(hari, " Hari")

}