package main

import "fmt"

func main() {
	var bilanganX int
	var bilanganY int
	fmt.Scan(&bilanganX, &bilanganY)
	for done := false; !done; {
		bilanganX = bilanganX - bilanganY
		fmt.Println(bilanganX)

		done = bilanganX <= 0
			
	}

	fmt.Println(bilanganX == 0)

}