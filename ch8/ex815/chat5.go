// “Failure of any client program to read data in a timely manner ultimately causes all clients to get stuck. Modify the broadcaster to skip a message rather than wait if a client writer is not ready to accept it. Alternatively, add buffering to each client’s outgoing message channel so that most messages are not dropped; the broadcaster should use a non-blocking send to this channel”

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
	// Define a bufferred channel
	bOut := make(chan<- string, 10)
	// Copied from the client channel
	bCli := client{bOut, ""}
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			// NOTE: Non-blocking sending message to client outgoing message channel
			for bCli = range clients {
				// DEBUG: Cap is shown as 10. Which might be working
				// DEBUG It its still 0 length means, unbufferred channel
				// fmt.Printf("current client message length: %d", len(bCli.Out))
				// fmt.Printf("current client message capacity: %d", cap(bCli.Out))
				select {
				case bCli.Out <- msg:
				default:
					fmt.Println("Channel is full, skip the message")
				}
			}

		case bCli = <-entering:
			clients[bCli] = true
			var onlines []string
			for c := range clients {
				onlines = append(onlines, c.Name)
			}
			bCli.Out <- fmt.Sprintf("%d clients: %s", len(clients), strings.Join(onlines, ", "))

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.Out)
		}
	}
}

//!-broadcaster

//!+handleConn communicate to a single client
func handleConn(conn net.Conn) {
	// Set timer to close the connection
	timeout := 5 * time.Second
	timer := time.NewTimer(timeout)

	go func() {
		<-timer.C
		conn.Close()
	}()

	// Force to input the client name first
	var who string
	fmt.Fprint(conn, "Please input your name: ")
	input := bufio.NewScanner(conn)
	input.Scan()
	who = input.Text()

	// Buffer 10 messages for client for testing
	// TODO: need to verify the blocking case if we send the 3rd msg
	// to the channel
	ch := make(chan string, 10) // outgoing client messages
	go clientWriter(conn, ch)

	cli := client{ch, who}
	cli.Out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	// input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(timeout)
	}
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
