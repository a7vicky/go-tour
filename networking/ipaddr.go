/*
type IPAddr {
	IP IP
}
A primary use of this type is to perform DNS lookup on IP hostname
func ResolveIPAddr(net, addr string) (*IPAddr, os.Error)
net is one of "ip", "ip4" or "ip6"
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

	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error)
		os.Exit(1)
	}
	fmt.Println("Resolution Address is ", addr.String())
	os.Exit(0)
}
