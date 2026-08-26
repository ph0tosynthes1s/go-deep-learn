package main

import (
	"fmt"
	"net"
	"os"
)

func lookIP(address string) ([]string, error) {
	hosts, err := net.LookupAddr(address)
	if err != nil {
		return nil, err
	}
	return hosts, nil
}

func lookHostname(hostname string) ([]string, error) {
	IPs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	return IPs, nil
}

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Pls provide an argument!")
		return
	}

	input := args[1]
	IPaddress := net.ParseIP(input)
	if IPaddress == nil {
		IPs, err := lookHostname(input)
		if err == nil {
			for _, ip := range IPs {
				fmt.Println(ip)
			}
		}
	} else {
		hosts, err := lookIP(input)
		if err == nil {
			for _, host := range hosts {
				fmt.Println(host)
			}
		}
	}
}
