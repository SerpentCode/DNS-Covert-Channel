package main

import (
	"CC/pkg/client"
	"flag"
)

func main() {

	sender := flag.Bool("sender", false, "Whether to open in sender mode")
	fileIn := flag.String("File In", "", "Filepath of file to send (sender mode only)")
	fileOut := flag.String("File Out", "", "Location to save received data")
	dnsIP := flag.String("DNS IP", "", "IP of DNS server to use")
	flag.Parse()

	if *sender {
		client.StartSender(*fileIn, *dnsIP)
	} else {
		client.StartReceiver(*fileOut)
	}
}
