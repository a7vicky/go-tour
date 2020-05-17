package main

import "fmt"

const a = 42
const b = 43.8
const c = "My name is Vicky"

func main() {
	fmt.Printf("Type of %v is %T\n", a, a)
	fmt.Printf("Type of %v is %T\n", b, b)
	fmt.Printf("Type of %v is %T\n", c, c)
}
