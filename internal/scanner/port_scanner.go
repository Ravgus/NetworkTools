package scanner

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func isPortOpen(host string, port int) bool {
	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort(host, strconv.Itoa(port)),
		3*time.Second,
	)

	if err != nil {
		return false
	}

	defer conn.Close()

	return true
}

func ScanPort(host string, port int, ch chan string) {
	if isPortOpen(host, port) {
		ch <- fmt.Sprintf("Host: %v - Port %v is open!", host, port)
	} else {
		ch <- fmt.Sprintf("Host: %v - Port %v is down!", host, port)
	}
}
