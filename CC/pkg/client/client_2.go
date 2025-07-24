// client.go
package client

import (
	"fmt"
	"log"
	"time"

	"github.com/miekg/dns"
)

func StartClient2() {
	client := new(dns.Client)
	domain := "www1.bobmovies.us."
	serverAddr := "192.168.13.31:53"

	// first query populates the cache
	warm := new(dns.Msg)
	warm.SetQuestion(domain, dns.TypeA)
	if _, _, err := client.Exchange(warm, serverAddr); err != nil {
		log.Fatalf("Warm-up query failed: %v", err)
	}

	// Actual measurement
	msg := new(dns.Msg)
	msg.SetQuestion(domain, dns.TypeA)

	start := time.Now()
	resp, _, err := client.Exchange(msg, serverAddr)
	duration := time.Since(start)

	if err != nil {
		log.Fatalf("DNS query failed: %v", err)
	}

	fmt.Printf("Query for %s took %v", domain, duration)

	// Infer cache hit/miss by threshold
	threshold := 1 * time.Second
	if duration < threshold {
		fmt.Println("✔ Cache hit detected")
	} else {
		fmt.Println("✘ Cache miss detected")
	}

	for _, ans := range resp.Answer {
		fmt.Println(ans.String())
	}
}
