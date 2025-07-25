package main

import (
	"CC/pkg/client"
	"time"
)

func main() {
	// server.StartDNS()
	// client.StartClient2()
	// sleep()
	client.StartSender()
	time.Sleep(time.Second * 3)
	client.StartReceiver()
}
