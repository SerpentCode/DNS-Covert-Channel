// client.go
package client

import (
	"fmt"
	"strconv"
	"time"

	"github.com/miekg/dns"
)

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

func StartClient1() {
	domain := "www1.bobmovies.us"
	serverAddr := "192.168.13.31:53"

	before, err := getInsertions(serverAddr)

	msg := new(dns.Msg)
	msg.SetQuestion(dns.Fqdn(domain), dns.TypeA)
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
		fmt.Println("✔ Cache hit detected")
	} else {
		fmt.Println("✘ Cache miss detected")
	}
}
