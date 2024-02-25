package main

import (
	"fmt"
	"net"
	"os"
)

func run() {
	// Redis is 6379 originally
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Failed to listen on port 6379...")
		os.Exit(1)
		return
	}
	fmt.Println("Listening on port 6379...")
	conn, err := ln.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	for {

		handle(conn)
		conn.Write([]byte("+OK\r\n"))
	}
}
