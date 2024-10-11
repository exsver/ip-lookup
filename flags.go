package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	Net *string
	IP  *string
}

func (f *Flags) String() string {
	str := ""

	if f.Net != nil {
		str += fmt.Sprintf("net:'%s'", *f.Net)
	}

	if f.IP != nil {
		str += fmt.Sprintf("ip:'%s'", *f.IP)
	}

	return str
}

func parseFlags() *Flags {
	var f Flags

	f.Net = flag.String("net", "0.0.0.0/0", "network in CIDR format")
	f.IP = flag.String("ip", "0.0.0.0", "ip address")

	flag.Parse()

	return &f
}
