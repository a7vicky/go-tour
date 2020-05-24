/* IPMask is of type byte slice and used for handling masking operations
type IPMask []byte
Function to create a mask from IPv4 address
func IPv4Mask(a,b,c,d byte) IPMask OR func (ip IP) DefaultMask() IPMask
*/

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
	dotAddr := os.Args[1]

	addr := net.ParseIP(dotAddr)
	if addr == nil {
		fmt.Println("Invalid Address")
		os.Exit(1)
	}
	// DefaultMask returns the default IP mask for the IP address ip. Only IPv4 addresses have default masks;
	// DefaultMask returns nil if ip is not a valid IPv4 address.

	mask := addr.DefaultMask()
	network := addr.Mask(mask)
	ones, bits := mask.Size()
	fmt.Println("Address is ", addr.String(),
		"DefaultMask length is ", bits,
		"Leading one count is ", ones,
		"Mask is hex ", mask.String(),
		" Network is ", network.String(),
	)
	os.Exit(0)
}
