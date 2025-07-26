package main

import (
	"CC/pkg/client"
	"flag"
)

func main() {

	fileIn := flag.String("file-in", "", "Filepath of file to send (sender mode only)")
	fileOut := flag.String("file-out", "", "Location to save received data")
	dnsIP := flag.String("dns", "", "IP of DNS server to use")
	flag.Parse()

	if *fileIn != "" {
		println("Going to sender")
		client.StartSender(*fileIn, *dnsIP)
	} else {
		println("Going to receiver")
		client.StartReceiver(*fileOut)
	}
}
