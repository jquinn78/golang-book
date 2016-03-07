package main

import "fmt"

func main() {

	const meters float64 = 0.3048

	fmt.Println("Enter the number of feet: ")

	var feetInput float64

	fmt.Scanf("%f", &feetInput)

	feetOutput := feetInput * meters

	fmt.Println("Meters: ", feetOutput)

}
