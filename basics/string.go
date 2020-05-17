package main

import (
	"fmt"
)

func main() {
	s := "Hello, Vicky"
	fmt.Println(s)
	fmt.Printf("Type of s is %T\n", s)
	// string is slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("Type of bs is %T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}
	for i := 0; i < len(s); i++ {
		fmt.Printf("%b \t", s[i])

	}
}
