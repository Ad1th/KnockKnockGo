package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func scanPort(host string, port int, openCount *int32, wg *sync.WaitGroup) {
	defer wg.Done()

	address := fmt.Sprintf("%s:%d", host, port)

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return // skip closed ports (clean output)
	}
	conn.Close()

	fmt.Println("🚪 Knock... Port", port, "is OPEN")
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