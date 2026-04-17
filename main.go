package main

import (
	"flag"
	"fmt"
	"net"
	"runtime"
	"sync"
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

type scanResult struct {
	port     int
	duration time.Duration
	service  string
}

func scanPort(host string, port int, timeout time.Duration) (scanResult, bool) {
	result := scanResult{port: port}

	address := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()

	conn, err := net.DialTimeout("tcp", address, timeout)
	result.duration = time.Since(start)
	if err != nil {
		return result, false
	}
	conn.Close()

	if service, ok := commonServices[port]; ok {
		result.service = service
	}

	return result, true
}

func worker(host string, timeout time.Duration, jobs <-chan int, results chan<- scanResult, wg *sync.WaitGroup) {
	defer wg.Done()

	for port := range jobs {
		if result, ok := scanPort(host, port, timeout); ok {
			results <- result
		}
	}
}

func printOpenResult(result scanResult) {
	if result.service != "" {
		fmt.Printf("🚪 Knock... Port %d is OPEN in %s (%s)\n", result.port, result.duration, result.service)
		if result.port == 80 {
			fmt.Println("HTTP server detected")
		}
		return
	}

	fmt.Printf("🚪 Knock... Port %d is OPEN in %s\n", result.port, result.duration)
}

func main() {
	server := flag.String("server", "localhost", "target host to scan")
	startPort := flag.Int("start", 1, "start port")
	endPort := flag.Int("end", 65535, "end port")
	workers := flag.Int("workers", runtime.NumCPU()*8, "number of concurrent workers")
	timeoutMS := flag.Int("timeout-ms", 500, "dial timeout in milliseconds")
	flag.Parse()

	if *startPort < 1 || *startPort > 65535 {
		fmt.Println("Invalid -start value: must be between 1 and 65535")
		return
	}

	if *endPort < 1 || *endPort > 65535 {
		fmt.Println("Invalid -end value: must be between 1 and 65535")
		return
	}

	if *startPort > *endPort {
		fmt.Println("Invalid range: -start cannot be greater than -end")
		return
	}

	if *workers < 1 {
		fmt.Println("Invalid -workers value: must be at least 1")
		return
	}

	if *timeoutMS < 1 {
		fmt.Println("Invalid -timeout-ms value: must be at least 1")
		return
	}

	fmt.Println("🔍 KnockKnockGo scanning...\n")

	host := *server
	portCount := *endPort - *startPort + 1
	if *workers > portCount {
		*workers = portCount
	}

	fmt.Printf("Target host: %s\n", host)
	fmt.Printf("Port range: %d-%d\n", *startPort, *endPort)
	fmt.Printf("Workers: %d\n", *workers)
	fmt.Printf("Timeout: %dms\n\n", *timeoutMS)

	jobs := make(chan int, *workers*2)
	results := make(chan scanResult, *workers*2)

	var workerWg sync.WaitGroup
	for i := 0; i < *workers; i++ {
		workerWg.Add(1)
		go worker(host, time.Duration(*timeoutMS)*time.Millisecond, jobs, results, &workerWg)
	}

	go func() {
		for port := *startPort; port <= *endPort; port++ {
			jobs <- port
		}
		close(jobs)
	}()

	go func() {
		workerWg.Wait()
		close(results)
	}()

	openCount := 0
	for result := range results {
		printOpenResult(result)
		openCount++
	}

	fmt.Println("Open ports:", openCount)
	fmt.Println("\n✅ Scan complete")
}