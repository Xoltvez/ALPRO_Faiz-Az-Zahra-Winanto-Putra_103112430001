package main

import "fmt"

func main() {

	var x, j, bilangan, nX, nZero int

	fmt.Scan(&x)

	for j = 1; j <= 9; j++ {
		fmt.Scan(&bilangan)

		if bilangan == x {
			nX++
		} else {
			nZero++
		}
	}

	if nX > nZero {
		fmt.Print("Modus = ", x)
	} else {
		fmt.Print("Modus = ", nZero)
	}

}