package main

import "fmt"

func ReturnHalf(number int) (int, bool) {

	halfNum := number / 2

	if halfNum%2 == 0 {
		return halfNum, true
	} else {
		return halfNum, false
	}

}

func main() {

	var number int
	fmt.Println("Enter the number to half: ")
	fmt.Scanf("%d", &number)

	fmt.Println(ReturnHalf(number))

}
