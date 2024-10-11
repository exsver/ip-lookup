package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	// parse flags
	flags := parseFlags()

	_, ipNet, err := net.ParseCIDR(*flags.Net)
	if err != nil {
		log.Fatal(err)
	}

	ip := net.ParseIP(*flags.IP)
	if ip == nil {
		log.Fatal("invalid IP address")
	}

	isContains := ipNet.Contains(ip)

	fmt.Println(isContains)

	if isContains {
		os.Exit(100) // Exit code 100 if true
	} else {
		os.Exit(200) // Exit code 200 if false
	}
}
