/* https://golang.org/ref/spec#Label_scopes
   https://golang.org/ref/spec#Labeled_statements
*/

package main

import "fmt"

func main() {
	var x int = 0

	// for loop work as a while loop
Lable1:
	for x < 8 {
		if x == 5 {

			// using goto statement
			x = x + 1
			goto Lable1
		}
		fmt.Printf("value is: %d\n", x)
		x++
	}
}
