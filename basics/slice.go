// https://blog.golang.org/slices-intro
package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4}
	fmt.Printf("Type of x is %T\n", x)
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	//
	var y = []int{5, 6, 7}
	fmt.Println(y)
	// using len function for getting the size of an array
	fmt.Println(len(y))
	fmt.Printf("Type of y is %T\n", y)
	// Variadic parameter to append
	// https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	x = append(x, y...)
	fmt.Println(x)

	// Range clause
	for i, v := range x {
		fmt.Printf("Value at index %v is %v \n", i, v)
	}
}
