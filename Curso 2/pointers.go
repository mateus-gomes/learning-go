package main

import "fmt"

func main() {
	x := 5
	// Creating a pointer
	y := &x
	// Changing the original value
	*y = 10

	fmt.Println(x, *y)
	fmt.Println(&x, y)
	printValues(&x)
	fmt.Println(x, *y)
	fmt.Println(&x, y)
}

func printValues(x *int) {
	fmt.Println("====Updating value====")
	*x = 20
}
