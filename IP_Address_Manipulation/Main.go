package main

import (
	"fmt"
	"net"
)

func main() {
	netw := "1.1.1.0/24"
	ips := []string{"1.1.1.1", "2.2.2.2"}

	for _, ip := range ips {
		_, cidrnet, _ := net.ParseCIDR(netw)

		addr := net.ParseIP(ip)
		result := cidrnet.Contains(addr)

		fmt.Printf("%s contains %s: %t \n", cidrnet, addr, result)

	}

	// var counter int8
	counter := iota

	for start, end := 0, 10; start < end; start = start + 1 {
		println(counter)
	}
}
