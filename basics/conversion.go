package main

import (
	"fmt"
)

type cluster string
type k8s string

func main() {
	var name cluster = "mycluster"
	fmt.Println(name)
	fmt.Printf("Type of name is %T\n", name)
	//	name = k8s(string(name)) is wrong as go is static typed language not dynamic typed like python
	var z = k8s(string(name))
	fmt.Printf("Type of name is %T", z)
}
