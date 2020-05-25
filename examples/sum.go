// func (r receiver) identifier(parameter(s)) (return(s)) { code}

package main

import "fmt"

func main() {
	xi := []int{2, 4, 5, 3, 5, 6, 3, 5, 5}
	fmt.Println(sum(xi...))
}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n", ii)
	total := 0
	for _, v := range ii {
		total += v
	}
	return total
}
