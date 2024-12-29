package main

import "fmt"

func main(){
	var x int

	fmt.Scan(&x)

	for i := 1; i <= x; i+= 1{
		if x % i == 0 {
			fmt.Print(i, "  ")
		}
}
}