package chapter8

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func StartEchoServer() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return // connection aborted
		}

		// Pass a connection to the handler
		go handleFunc(conn)
	}
}

func handleFunc(conn net.Conn) {
	// Extract the input from the connection
	input := bufio.NewScanner(conn)

	// Scan the content
	for input.Scan() {
		echo(conn, input.Text(), 1*time.Second)
	}
	conn.Close()
}

func echo(conn net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(shout))
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t", shout)
	time.Sleep(delay)

	fmt.Fprintln(conn, "\t", strings.ToLower(shout))
	time.Sleep(delay)
}
