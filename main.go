package main

import (
	"fmt"
	"net"
	"time"

	"github.com/racingmars/go3270"
)

func main() {
	fmt.Println("Starting TN3270 server on :3270")
	ln, err := net.Listen("tcp", ":3270")
	if err != nil {
		panic(err)
	}
	defer ln.Close()
	fmt.Println("Listening on :3270")

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
		time.Sleep(1 * time.Second)
	}
}