package main

import "fmt"

func main() {
	m := map[string]float64{
		"vicky": 5.7,
	}
	// https://golang.org/ref/spec#Deletion_of_map_elements
	delete(m, "test") // No such key exists
	for k, v := range m {
		println(k, v)
	}

	if v, ok := m["test"]; ok {
		fmt.Println("value:", v)
		delete(m, "test")
	}
	if v, ok := m["vicky"]; ok {
		fmt.Println(ok)
		fmt.Println("value:", v)
		delete(m, "vicky")
	}
	fmt.Println(m)
}
