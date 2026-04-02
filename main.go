package main

import (
	"flag"
	"fmt"
	"net"
	"time"

	"github.com/racingmars/go3270"
)

func main() {
	tn3270Port := flag.Int("port", 3270, "Port on which TN3270 server is served")
	flag.Parse()
	
	fmt.Printf("Starting TN3270 server on port %d\n", *tn3270Port)
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", *tn3270Port))
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Server listening")

	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("New connection from %s\n", conn.RemoteAddr())
	defer conn.Close()

	_, err := go3270.NegotiateTelnet(conn)
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}

	// Infinite loop to keep connection open
	for {
		conditions := getConditions()
		scr := toScreen(conditions)
		go3270.ShowScreenOpts(scr, nil, conn, go3270.ScreenOpts{
			NoResponse: true, // do not block on waiting for user's input
		})
		// attempt to refresh every minute (note that we already cache data for 1 hour)
		time.Sleep(1 * time.Minute)
	}
}