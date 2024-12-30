package main

import "fmt"

func main() {

	var j, y, bilangan, total int

	fmt.Scan(&y)

	total = 0

	for j = 1; j <= 9; j++ {
		fmt.Scan(&bilangan)
		total += bilangan
	}

	if total >= y*5 {
		fmt.Print("Median Bernilai ", y)
	} else {
		fmt.Print("Median Bernilai 0")
	}

}