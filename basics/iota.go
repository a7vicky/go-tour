package main

import "fmt"

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d \t\t\t%b\n", kb, kb)
	fmt.Printf("%d \t\t%b\n", mb, mb)
	fmt.Printf("%d \t\t%b", gb, gb)

	foo()
	bar()
}

func foo() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
func bar() {
	const (
		a = 1 << iota // a == 1  (iota == 0)
		b = 1 << iota // b == 2  (iota == 1)
		c = 3         // c == 3  (iota == 2, unused)
		d = 1 << iota // d == 8  (iota == 3)
	)

	const (
		u         = iota * 42 // u == 0     (untyped integer constant)
		v float64 = iota * 42 // v == 42.0  (float64 constant)
		w         = iota * 42 // w == 84    (untyped integer constant)
	)

	const x = iota // x == 0
	const y = iota // y == 0

	fmt.Println(a, b, c, d)
	fmt.Println(u, v, w)
	fmt.Println(x, y)
}
