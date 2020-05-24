/* In example ipaddr.go we look into the func ResolveIPAddr
However host may have mutiple IP addresses, usually in case of multiple NIC attached
and could have mutiple hostnames acting as alias
==> func LookupHost(name string) (addre []string, err os.Error)
One of  these addresses will be labelles as "canonical" hostname using below function
==> func LookupCNAME(name string) (cname strig, err os.Error)
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname \n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {
		fmt.Println(s)
	}
	os.Exit(0)
}
