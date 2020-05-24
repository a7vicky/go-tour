/* net pckage define many types
type IP is define as byte slice
type IP []byte
*/

//net package: https://golang.org/pkg/net/
//os package:  https://golang.org/pkg/os/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		// https://golang.org/pkg/fmt/#Fprintf
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid Address")
	} else {
		fmt.Println("Address is ", addr.String())
	}
	os.Exit(0)
}
