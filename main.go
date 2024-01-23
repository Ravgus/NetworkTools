package main

import (
	"bufio"
	"fmt"
	"github.com/Ravgus/NetworkTools/internal/loading"
	"github.com/Ravgus/NetworkTools/internal/scanner"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Select your goal: ")
	fmt.Println("1. Check available ports")
	fmt.Println("2. Check server availability")
	fmt.Println("3. DDOS")
	fmt.Print("Type an appropriate number: ")

	goal := getIntInput(reader)

	switch goal {
	case 1:
		scanPorts(reader)
	case 2:
		pingHost(reader)
	case 3:
		makeDDOS(reader)
	default:
		fmt.Println("Unsupported action!")

		return
	}
}

func scanPorts(reader *bufio.Reader) {
	fmt.Print("Enter host: ")

	host := getStringInput(reader)

	fmt.Print("Enter start port: ")

	startPort := getIntInput(reader)

	fmt.Print("Enter end port: ")

	endPort := getIntInput(reader)

	if startPort > endPort || startPort < 1 || endPort > 65535 {
		fmt.Println("Wrong port selected!")

		return
	}

	channel := make(chan string)

	for port := startPort; port <= endPort; port++ {
		go scanner.ScanPort(host, port, channel)
	}

	for port := startPort; port <= endPort; port++ {
		fmt.Println(<-channel)
	}
}

func makeDDOS(reader *bufio.Reader) {
	fmt.Print("Enter host: ")
	host := getStringInput(reader)

	fmt.Print("Enter a number of threads: ")
	threads := getIntInput(reader)

	loading.StartDDOS(host, threads)
}

func pingHost(reader *bufio.Reader) {
	fmt.Print("Enter host (with corresponded protocol): ")
	host := getStringInput(reader)

	scanner.PingHost(host)
}

func getStringInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	input = strings.TrimSpace(input)

	return input
}

func getIntInput(reader *bufio.Reader) int {
	port, err := strconv.Atoi(getStringInput(reader))

	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	return port
}
