/* https://tour.golang.org/flowcontrol/12
https://blog.golang.org/defer-panic-and-recover
A defer statement postpones the execution of a function until the surrounding function returns, either normally or through a panic.
Deferred calls are executed even when the function panics:
*/

package main

import (
	"fmt"
)

func main() {
	defer foo()
	bar()
	test()
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

func test() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
