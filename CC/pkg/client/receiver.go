// client.go
package client

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

var file *os.File
var serverAddr string
var DOMAIN_INDEX = 0

/*
Warm up DNS cache
*/
func init() {
	getInsertions(serverAddr)
}

/*
Helper: Converts an array of bits (typed as a byte slice) to an array of bytes
*/
func bitstoBytes(bits []byte) []byte {
	nBytes := (len(bits) + 7) / 8
	out := make([]byte, nBytes)

	for i, b := range bits {
		if b == 1 {
			byteIdx := i / 8
			bitIdx := 7 - uint(i%8)
			out[byteIdx] |= 1 << bitIdx
		}
	}
	// Return all but colon
	return out
}

/*
Helper: Determines if a dns query was cached
*/
func getInsertions(addr string) (int, error) {
	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn("insertions.bind"), dns.TypeTXT)
	msg.Question[0].Qclass = dns.ClassCHAOS
	c := &dns.Client{Timeout: 2 * time.Second}

	resp, _, err := c.Exchange(msg, addr)

	if err != nil {
		return 0, err
	}

	if len(resp.Answer) == 0 {
		return 0, fmt.Errorf("no insertion.bind response")
	}

	txt := resp.Answer[0].(*dns.TXT).Txt
	if len(txt) == 0 {
		return 0, fmt.Errorf("empty insertions.bind TXT")
	}
	return strconv.Atoi(txt[0])
}

func StartReceiver(filepath, dnsIP string) {
	serverAddr = dnsIP + ":53"
	// Parse size from header
	size_s := getSize()

	size, _ := strconv.Atoi(size_s)
	data_bits := make([]byte, 0)

	// Use the size to pull data from the DNS server by querying the cache
	for range size * 8 {
		data_bits = append(data_bits, readBit())

	}
	// We receive data as 'bits' displayed in a 'byte' format, so 00000000 or 00000001, so we convert this to an array
	// of 'bits', 1 or 0, in a byte slice
	data_bytes := bitstoBytes(data_bits)
	// Now we convert our byte slice of 'bits' to actual bytes in a byte slice
	this_shit_weird := bitstoBytes(data_bytes)
	// Write our data to file
	fmt.Println("Finished receiving data")
	if err := os.WriteFile(filepath, this_shit_weird, 0o644); err != nil {
		fmt.Println("Write file error: check file path?")
	}
}

/*
Gets size of the received file, located in the header
*/
func getSize() string {
	pat := []byte{0, 0, 1, 1, 1, 0, 1, 0}
	window := make([]byte, 0, len(pat))
	var prefix []byte
	for {
		bit := readBit()

		prefix = append(prefix, bit)

		window = append(window, bit)

		if len(window) > len(pat) {
			window = window[1:]
		}
		if bytes.Equal(window, pat) {
			// God I hate this fucking byte array but they are bits shit, NOT A FUNNY BIT
			prefix_bytes := bitstoBytes(prefix)
			prefix_bytes = prefix_bytes[:len(prefix_bytes)-1]
			return string(prefix_bytes)
		}
	}
}

// Queries the chosen DNS and determines if the dns query is cached
func readBit() byte {
	before, err := getInsertions(serverAddr)

	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(DomainList[DOMAIN_INDEX]), dns.TypeA)
	c := &dns.Client{Timeout: 2 * time.Second}
	_, _, err = c.Exchange(msg, serverAddr)

	if err != nil {
		fmt.Println("✘ Cache miss really detected for " + DomainList[DOMAIN_INDEX] + " consider increasing timeout " + err.Error())
	}

	after, err := getInsertions(serverAddr)
	// if err != nil {
	// 	panic(err)
	// }

	if before == after {
		fmt.Println("\033[32m✔ Cache hit detected for " + DomainList[DOMAIN_INDEX] + "\033[0m")
		DOMAIN_INDEX++
		return 1
	} else {
		fmt.Printf("\033[31m✘ Cache miss detected for %s\n\033[0m", DomainList[DOMAIN_INDEX])
		DOMAIN_INDEX++
		return 0
	}
}
