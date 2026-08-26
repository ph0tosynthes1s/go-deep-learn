package main

import (
	"fmt"
	"log"
	"net"
)

func GetOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr()
	localIp := localAddr.(*net.UDPAddr).IP
	return localIp
}

func main() {
	ip := GetOutboundIP()
	fmt.Println("Your outbound IP address is:", ip.String())
}
