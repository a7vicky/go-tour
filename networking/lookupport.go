/* LookupPort function used to get port from the /etc/services file in Unix system
==> func LookupPort(network, service string) (port int, err os.Error)
example network argument is "tcp" or "udp" and service such as "ssh"
*/

package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s network-type service \n", os.Args[0])
		os.Exit(1)
	}
	networkType := os.Args[1]
	service := os.Args[2]

	port, err := net.LookupPort(networkType, service)

	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	fmt.Println("Service port", port)
	os.Exit(0)
}
