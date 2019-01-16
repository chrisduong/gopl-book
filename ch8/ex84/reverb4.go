// “Modify the reverb2 server to use a sync.WaitGroup per connection to count the number of active echo goroutines. When it falls to zero, close the write half of the TCP connection as described in Exercise 8.3. Verify that your modified netcat3 client from that exercise waits for the final echoes of multiple concurrent shouts, even after the standard input has been closed”

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}

//!+
func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	fmt.Println("Im here 1!")
	fmt.Println(input.Scan())
	for input.Scan() {
		fmt.Println("Im here 2!")
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	wg.Wait()
	// NOTE: ignoring potential errors from input.Err()
	// NOTE: convert conn to TCP connection so we can do CloseWrite
	if tcpCon, ok := c.(*net.TCPConn); ok {
		_ = tcpCon.CloseWrite() // Ignore errors
	}
}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
