package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// obtained from ping -c 1 stackoverflow.com, should print "stackoverflow.com"
	addr, err := net.LookupAddr(os.Args[1])
	if err != nil {
		
	}

}
