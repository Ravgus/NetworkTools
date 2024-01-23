package scanner

import (
	"fmt"
	"net/http"
)

func ping(host string) bool {
	_, err := http.Get(host)

	return err != nil
}

func PingHost(host string) {
	if ping(host) {
		fmt.Printf("Host: %v is up!\n", host)
	} else {
		fmt.Printf("Host: %v is down!\n", host)
	}
}
