// “Make the chat server disconnect idle clients, such as those that have sent no messages in the last five minutes. Hint: calling conn.Close() in another goroutine unblocks active Read calls such as the one done by input.Scan()”

// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster

// client contain its Name and its outgoing message channel
type client struct {
	Out  chan<- string // an outgoing message channel
	Name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

// broadcast broadcasts messages to all client
func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli.Out <- msg
			}

		case cli := <-entering:
			cli.Out <- "Current clients: "
			for c := range clients {
				cli.Out <- c.Name
			}
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

//!-broadcaster

//!+handleConn communicate to a single client
func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	// NOTE: type coercion `ch` to `client` object so you can pass it to
	// channel `entering`
	cli := client{ch, who}
	cli.Out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	timeout := 4 * time.Second
	timer := time.NewTimer(timeout)

	go func() {
		<-timer.C
		conn.Close()
	}()

	input := bufio.NewScanner(conn)
	// Make new goroutine here
	go func() {
		for input.Scan() {
			messages <- who + ": " + input.Text()
			timer.Reset(timeout)
		}
	}()
	// NOTE: ignoring potential errors from input.Err()
	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

//!+main
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

//!-main

