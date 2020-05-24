package main

func main() {
	m := map[string]float64{
		"vicky": 5.7,
		"tom":   6.2,
		"harry": 5.9,
	}
	// https://golang.org/ref/spec#For_statements
	for k, v := range m {
		println(k, v)
	}
}
