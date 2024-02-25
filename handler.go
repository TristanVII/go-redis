package main

import (
	"bufio"
	"fmt"
	"io"
)

func handle(rd io.Reader) {
	reader := bufio.NewReader(rd)

	b, _ := reader.ReadString('\n')
	fmt.Println("Received: " + string(b))

}
