package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func getIP(fqdn string) []net.IP {
	ip, err := net.LookupIP(fqdn)
	if err != nil {
		fmt.Println(err)
	}
	return ip
}

func getHost(ip string) []string {
	host, err := net.LookupHost(ip)
	if err != nil {
		fmt.Println(err)
	}
	return host
}

func checkInput(checkstr string) string {
	addr := net.ParseIP(checkstr)
	if addr != nil {
		result, _ := net.LookupAddr(checkstr)
		return fmt.Sprintf("%s\n", result)
	} else {
		result, _ := net.LookupIP(checkstr)
		return fmt.Sprintf("%s\n", result)
	}
}

func cmdMode() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("#>: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("something went wrong getting input")
		}
		text = strings.Replace(text, "\n", "", -1)
		if text == "exit" {
			os.Exit(1)
		}
		addr := net.ParseIP(text)
		if addr != nil {
			result, _ := net.LookupAddr(text)
			fmt.Printf("%s\n", result)
		} else {
			result, _ := net.LookupIP(text)
			fmt.Printf("%s\n", result)
		}
	}
}

func main() {
	// check if any arguments passed
	if len(os.Args[1:]) == 0 {
		// no args = cmdMode
		cmdMode()
	}
	// check what the argument is
	argument := os.Args[1]
	// is arg is a file, resolve each line
	if _, err := os.Stat(argument); err == nil {
		file, err := os.Open(argument)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(getIP(scanner.Text()))
		}
		// if arg is not a file, check it and resolve it.
	} else if os.IsNotExist(err) {
		fmt.Println(checkInput(argument))

	} else {
		fmt.Println(getHost(argument))
	}

}
