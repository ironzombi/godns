package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
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
		return fmt.Sprintf("%s", result)
	} else {
		result, _ := net.LookupIP(checkstr)
		return fmt.Sprintf("%s", result)
	}
}

func cmdMode() {

}

func main() {
	if len(os.Args[1:]) == 0 {
		cmdMode()
	}
	argument := os.Args[1]
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

	} else if os.IsNotExist(err) {
		fmt.Println(checkInput(argument))

	} else {
		fmt.Println("um, try again ?")
	}

}
