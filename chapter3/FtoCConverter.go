package main

import "fmt"

func main() {

	fmt.Printf("Enter the temperature (F) to convert: ")
	var farenheight float64
	fmt.Scanf("%f", &farenheight)

	celcius := (farenheight - 32) * 5 / 9

	fmt.Println("Celcisus: ", celcius)

}
