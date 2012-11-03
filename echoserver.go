/*
 Noir Chat IRC Server
*/

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	RPL_WELCOME  = 1
	)

func main() {
	// Listen on TCP port 2000 on all interfaces.
	fmt.Println("Starting up on port 6667");
	l, err := net.Listen("tcp", ":6667")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Wait for a connection.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
			// multiple connections may be served concurrently.
		//go echoConnection(conn)
			go echoWithPrefix(conn)
	}
}

// this uses the standard library io.Copy
// method to implement an echo server.
// Flow of control is a little less clear
// because Copy just continues to read
// internally;
func echoConnection(c net.Conn) {
	// Echo all incoming data.
	io.Copy(c, c)
	// Shut down the connection.
	c.Close()
}

// Here I just use the Read interface
// on the connection object and so
// I have to manage continuing to read
// explicitly with the for(ever) loop.
func echoWithPrefix(c net.Conn) {
	buf := make( []byte, 1024 )
	for {
		_, err := c.Read(buf)
		if err != nil  {
			c.Close()
		}
		response := bytes.Join([][]byte{[]byte("server:"), buf}, []byte(" "))
		_, err1 := c.Write(response)
		if err1 != nil {
			c.Close()
		}
	}
}