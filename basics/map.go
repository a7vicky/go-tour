// https://golang.org/ref/spec#Map_types

package main

import "fmt"

func main() {
	m := map[string]int{
		"vicks":    26,
		"abhishek": 27,
		"tom":      34,
	}
	fmt.Println(m)
	fmt.Println(m["vicks"])
	v, ok := fmt.Println(m["dan"])
	fmt.Printf("%v and %v\n", v, ok)
	// fmt.Println(v)
	r, err := m["tom"] // err will be "true" if the key is valid
	fmt.Printf("%T\n", err)
	if !err {
		fmt.Println("there is an error : No key found")
	}
	fmt.Println(r)
}
