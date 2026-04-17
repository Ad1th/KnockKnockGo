package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

var commonServices = map[int]string{
	20:   "FTP Data",
	21:   "FTP Control",
	22:   "SSH",
	23:   "Telnet",
	25:   "SMTP",
	53:   "DNS",
	67:   "DHCP Server",
	68:   "DHCP Client",
	69:   "TFTP",
	80:   "HTTP",
	110:  "POP3",
	111:  "RPCBind",
	119:  "NNTP",
	123:  "NTP",
	135:  "MSRPC",
	137:  "NetBIOS Name",
	138:  "NetBIOS Datagram",
	139:  "NetBIOS Session",
	143:  "IMAP",
	161:  "SNMP",
	162:  "SNMP Trap",
	179:  "BGP",
	194:  "IRC",
	389:  "LDAP",
	443:  "HTTPS",
	445:  "SMB",
	465:  "SMTPS",
	514:  "Syslog",
	515:  "LPD",
	587:  "SMTP Submission",
	631:  "IPP",
	636:  "LDAPS",
	873:  "rsync",
	993:  "IMAPS",
	995:  "POP3S",
	1025: "NFS/Windows RPC",
	1080: "SOCKS Proxy",
	1194: "OpenVPN",
	1433: "MSSQL",
	1521: "Oracle DB",
	1723: "PPTP",
	1883: "MQTT",
	2049: "NFS",
	2375: "Docker",
	2376: "Docker TLS",
	3000: "Node/Web Dev",
	3128: "Squid Proxy",
	3306: "MySQL",
	3389: "RDP",
	3478: "STUN",
	3690: "SVN",
	4369: "Erlang EPMD",
	5000: "Flask/UPnP",
	5432: "PostgreSQL",
	5672: "AMQP (RabbitMQ)",
	5900: "VNC",
	5984: "CouchDB",
	6379: "Redis",
	6443: "Kubernetes API",
	6667: "IRC",
	7001: "WebLogic",
	8080: "HTTP Alt",
	8081: "HTTP Alt 2",
	8443: "HTTPS Alt",
	8888: "Jupyter",
	9042: "Cassandra",
	9092: "Kafka",
	9200: "Elasticsearch",
	9300: "Elasticsearch Cluster",
	9418: "Git",
	11211: "Memcached",
	15672: "RabbitMQ Admin",
	27017: "MongoDB",
	50000: "DB2",
}

func scanPort(host string, port int, openCount *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	duration := time.Since(start)
	if err != nil {
		return // skip closed ports (clean output)
	}
	conn.Close()

	if service, ok := commonServices[port]; ok {
		fmt.Printf("🚪 Knock... Port %d is OPEN in %s (%s)\n", port, duration, service)
		if port == 80 {
			fmt.Println("HTTP server detected")
		}
	} else {
		fmt.Printf("🚪 Knock... Port %d is OPEN in %s\n", port, duration)
	}

	atomic.AddInt32(openCount, 1)
}

func main() {
	server := flag.String("server", "localhost", "target host to scan")
	flag.Parse()

	fmt.Println("🔍 KnockKnockGo scanning...\n")

	host := *server
	fmt.Printf("Target host: %s\n\n", host)

	var wg sync.WaitGroup
	var openCount int32

	for port := 1; port <= 10000; port++ {
		wg.Add(1)
		go scanPort(host, port, &openCount, &wg)
	}

	wg.Wait()

	fmt.Println("Open ports:", openCount)
	fmt.Println("\n✅ Scan complete")
}