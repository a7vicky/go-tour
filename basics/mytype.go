package main

import "fmt"

var a int

// custom type k8s
type k8s string

// variable name of kind k8s
var name k8s

func main() {
	a = 42
	fmt.Println(a)
	fmt.Printf("%T \t %v \n", a, a)
	name = "kube8s"
	fmt.Println(name)
	var z k8s = "kubernetes"
	fmt.Println(z)
	fmt.Printf("Type of varibale z is %T\n", z)

	foo()
}

func foo() {
	fmt.Printf("Type of name is %T", name)
}
