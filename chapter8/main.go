package main

import "fmt"
import "golang-book/chapter8/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	max := math.Max(xs)
	min := math.Min(xs)
	fmt.Println(avg)
	fmt.Println(max)
	fmt.Println(min)
}
