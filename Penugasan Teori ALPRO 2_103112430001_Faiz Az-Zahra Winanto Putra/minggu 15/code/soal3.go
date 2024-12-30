package main

import "fmt"

func main() {

	var jam12, jam24 int

	fmt.Scan(&jam24)

	if jam24%12 == 0 {
		jam12 = 12
	} else {
		jam12 = jam24 % 12
	}

	if jam24 < 12 {
		fmt.Print(jam12, " AM")
	} else {
		fmt.Print(jam12, " PM")
	}

}