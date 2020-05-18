// https://golang.org/ref/spec#Array_types
// https://golang.org/ref/spec#Slice_types
package main

import "fmt"

func main() {

	// variable a which is array of integer type having lenth 5
	var a [5]int
	fmt.Println("emp:", a)

	// assign value at 4th index
	a[4] = 100
	fmt.Println("set:", a)
	// get the value of element at 4th index in the array
	fmt.Println("get:", a[4])

	// print the lengith of rray
	fmt.Println("len:", len(a))

	// Slice of an array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
