// https://gobyexample.com/slices
package main

import "fmt"

func main() {
	// https://tour.golang.org/moretypes/13
	// https://golang.org/ref/spec#Making_slices_maps_and_channels
	// https://golang.org/doc/effective_go.html#allocation_make
	s := make([]string, 3)
	n := make([]int, 2, 3)
	s[0] = "a"
	s[2] = "b"
	fmt.Println(s)
	s = append(s, "e", "f")
	fmt.Println(s)

	n[0] = 1
	n[1] = 2
	fmt.Println(n)
	fmt.Println(len(n)) // length of slice
	fmt.Println(cap(n)) // Capacity of slice
	// n[2] = 3
	//"error message: panic: runtime error: index out of range"
	n = append(n, 5, 6, 7, 8)
	fmt.Println(n)
	fmt.Println(len(n))
	fmt.Println(cap(n))
}
