package main

import "fmt"

// https://golang.org/ref/spec#Appending_and_copying_slices

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s)
	s = append(s, 9)
	fmt.Printf("new slice %v of type %T \n", s, s)
	n := []int{20, 23, 26}
	s = append(s, n...)
	fmt.Println(s)
}
