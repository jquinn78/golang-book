package main

import "fmt"

func GreatestNum(nums ...int) int {

	greatest := 0

	for _, value := range nums {

		if value > greatest {
			greatest = value
		}
	}

	return greatest
}

func main() {

	fmt.Println(GreatestNum(3, 100, 1, 1000, 2, 5, 101, 105, 10005))

}
