package main

import (
	"flag"
	"net"
	"zportscanner/zportscanner"
)

var i = flag.String("i", "127.0.0.1", "the ip address or ip file for port scanning")
var pr = flag.String("pr", "8000-9000", "the port range for scanning")

func main() {
	flag.Parse()
	if nil != net.ParseIP(*i) {
		zportscanner.ScanSingleIP(*i, *pr)
	} else {
		zportscanner.ScanIPFile(*i, *pr)
	}
}
