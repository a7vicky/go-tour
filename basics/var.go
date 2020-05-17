package main

import "fmt"

// https://golang.org/ref/spec#The_zero_value
var y int
var z int = 33

func main() {
	x := 12
	fmt.Println(x)
	foo()
	bar()
}

func foo() {
	fmt.Println(y)
	fmt.Println(z)
}

func bar() {
	type T struct {
		i    int
		f    float64
		next *T
	}
	t := new(T)
	fmt.Println(t)
}
