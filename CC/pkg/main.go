package main

import (
	"flag"
	"fmt"
	"os"
)

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func main() {

	filepath := flag.String("file", "", "File to exfiltrate")
	serverIP := flag.String("ip", "", "IP to send to")
	flag.Parse()

	if fileExists(*filepath) == false {
		fmt.Printf("File \"%s\" not found.\n", *filepath)
	}
	fmt.Println(*serverIP)

	//client.StartClient1()
}
