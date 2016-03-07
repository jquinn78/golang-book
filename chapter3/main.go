package main

import "fmt"

//this is a comment
/*this is also a comment
 */

func main() {
	var x string = "Hello World"
	var y string
	var z string
	y = "Hello World"
	fmt.Println(x)
	fmt.Println(y)
	z = "first"
	fmt.Println(z)
	z = "second"
	fmt.Println(z)
	z = z + "third"
	fmt.Println(z)
}
