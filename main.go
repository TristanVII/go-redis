package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	go run()
	conn, err := net.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Input and send loop
	reader := bufio.NewReader(os.Stdin) // Read from standard input
	for {
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text) // Remove newline

		if text == "exit" || text == "Exit" || text == "quit" || text == "Quit" || text == "stop" || text == "Stop" {
			fmt.Println("Exiting...")
			os.Exit(1)
		}

		_, err = fmt.Fprintln(conn, text)
		if err != nil {
			fmt.Println("Error sending:", err)
			return
		}
	}

}
