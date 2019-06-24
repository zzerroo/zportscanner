package zportscanner

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func checkFormat(ip, portRange string) (int, int, error) {
	portRanges := strings.Split(portRange, "-")
	if len(portRanges) != 1 && len(portRanges) != 2 {
		return -1, -1, fmt.Errorf("bad port ranges:%s", portRange)
	}

	startPort, erro := strconv.Atoi(portRanges[0])
	if erro != nil {
		return -1, -1, fmt.Errorf("bad port format:%s", portRanges[0])
	}

	endPort := startPort
	if len(portRanges) == 2 {
		endPort, erro = strconv.Atoi(portRanges[1])
		if erro != nil {
			return -1, -1, fmt.Errorf("bad port format:%s", portRanges[1])
		}
	}

	if startPort > endPort || startPort <= 0 || endPort >= 65536 {
		return -1, -1, fmt.Errorf("bad port ranges:%s", portRange)
	}

	if nil == net.ParseIP(ip) {
		return -1, -1, fmt.Errorf("bad ip format:%s", ip)
	}

	return startPort, endPort, nil
}

// ScanSingleIP scan the ip and port range to check
// wheather there is a process listen on the port
func ScanSingleIP(ip, portRange string) error {
	startPort, endPort, erro := checkFormat(ip, portRange)
	if erro != nil {
		log.Fatalf("%s", erro.Error())
		return erro
	}

	fmt.Printf("%s:\n", ip)
	for i := startPort; i <= endPort; i++ {
		conn, erro := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", ip, i), 300*time.Millisecond)
		if erro == nil {
			defer conn.Close()
			fmt.Println(fmt.Sprintf("\t%d", i))
		}
	}
	return nil
}

// ScanIPFile read the lines from ipFileName，to check wheather
// there is a process listen on the port
func ScanIPFile(ipFileName, portRange string) error {
	ipFile, erro := os.Open(ipFileName)
	if erro != nil {
		log.Fatalf("%s", erro.Error())
		return erro
	}

	input := bufio.NewScanner(ipFile)
	for input.Scan() {
		ScanSingleIP(input.Text(), portRange)
	}
	return nil
}
