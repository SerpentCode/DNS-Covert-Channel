// server.go
package server

import (
	"log"
	"sync"
	"time"

	"github.com/miekg/dns"
)

var (
	cache     = make(map[string]dns.RR)
	mu        sync.RWMutex
	lastSleep time.Duration
	upstream  = "8.8.8.8:53"
)

func handleDNS(w dns.ResponseWriter, r *dns.Msg) {
	q := r.Question[0]
	domain := q.Name

	mu.RLock()
	rr, ok := cache[domain]
	mu.RUnlock()

	var resp *dns.Msg

	if ok {
		// Cache hit: no sleep
		lastSleep = 0
		resp = new(dns.Msg)
		resp.SetReply(r)
		resp.Answer = []dns.RR{rr}
	} else {
		// Cache miss: simulate upstream delay
		sleepDuration := 2 * time.Second
		time.Sleep(sleepDuration)
		lastSleep = sleepDuration

		// Forward to real DNS
		c := new(dns.Client)
		in, _, err := c.Exchange(r, upstream)
		if err != nil {
			log.Printf("Upstream query error: %v", err)
			return
		}
		if len(in.Answer) > 0 {
			mu.Lock()
			cache[domain] = in.Answer[0]
			mu.Unlock()
		}
		resp = in
	}

	if err := w.WriteMsg(resp); err != nil {
		log.Printf("Failed to write msg: %v", err)
	}
}

func StartDNS() {

	// Start DNS server
	dns.HandleFunc(".", handleDNS)
	srv := &dns.Server{Addr: ":8053", Net: "udp"}
	log.Println("Mock DNS server listening on :8053")
	log.Fatal(srv.ListenAndServe())
}
