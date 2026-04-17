package main

import (
	"fmt"
	"net"
	"time"
)

func scanPort(port int) {
	address := fmt.Sprintf("localhost:%d", port)

	conn, err := net.DialTimeout("tcp", address, 300*time.Millisecond)
	if err != nil {
		return // skip closed ports (clean output)
	}
	conn.Close()

	fmt.Println("🚪 Knock... Port", port, "is OPEN")
}

func main() {
	fmt.Println("🔍 KnockKnockGo scanning...\n")

	for port := 1; port <= 10000; port++ {
		go scanPort(port)
	}

	time.Sleep(2 * time.Second)

	fmt.Println("\n✅ Scan complete")
}