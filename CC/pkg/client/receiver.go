// client.go
package client

import (
	"bytes"
	"fmt"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

// Determines if a dns query was cached
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

var serverAddr = "192.168.13.3:53"

func StartReceiver() {
	size := getSize()
	fmt.Println(size)
}

func getSize() string {
	const delim = "00111010"
	pat := []byte(delim)
	window := make([]byte, 0, len(pat))
	var prefix []byte
	i := 0
	for {
		bit := readBit(i)

		prefix = append(prefix, bit)

		window = append(window, bit)

		if len(window) > len(pat) {
			window = window[1:]
		}
		fmt.Printf("%b\n", window)
		if len(window) == len(pat) && bytes.Equal(window, pat) {
			return string(prefix)
		}

		i++
	}
}

// Queries the chosen DNS and determines if the dns query is cached
func readBit(index int) byte {
	before, err := getInsertions(serverAddr)

	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(DomainList[index]), dns.TypeA)
	c := &dns.Client{Timeout: 2 * time.Second}
	_, _, err = c.Exchange(msg, serverAddr)

	if err != nil {
		fmt.Println("✘ Cache miss really detected")
	}

	after, err := getInsertions(serverAddr)
	if err != nil {
		panic(err)
	}

	if before == after {
		fmt.Println("✔ Cache hit detected for " + DomainList[index])
		return 1
	} else {
		fmt.Println("✘ Cache miss detected for " + DomainList[index])
		return 0
	}
}
