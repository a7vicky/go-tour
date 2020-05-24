package main

import "fmt"

// https://golang.org/ref/spec#Struct_types

type person struct {
	first string
	last  string
}

func main() {
	p1 := person{
		first: "Abhishek",
		last:  "vicky",
	}
	p2 := person{
		first: "tod",
		last:  "mackey",
	}
	fmt.Println(p1, p2)
	fmt.Println(p1.first)
	fmt.Println(p2.last)
}
