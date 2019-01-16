// “Using a select statement, add a timeout to the echo server from Section 8.3 so that it disconnects any client that shouts nothing within 10 seconds.”
// We use `4 seconds` for faster debug

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
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))

}

//!+
func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	input := bufio.NewScanner(c)
	lines := make(chan string)

	defer func() {
		wg.Wait()
		if tcpCon, ok := c.(*net.TCPConn); ok {
			_ = tcpCon.CloseWrite()
		}
	}()

	go func() {
		for input.Scan() {
			lines <- input.Text()
		}
	}()

	// handle 4 seconds timeout
	timeout := 4 * time.Second
	timer := time.NewTimer(timeout)
	// Need to poll a channel to get any inputs
	for {
		select {
		case <-timer.C:
			return
			// need case for retrieving sth from the channel when reading
			// before sending to client.
		case line := <-lines:
			timer.Reset(timeout)
			wg.Add(1)
			go echo(c, line, 1*time.Second, &wg)
		}
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
