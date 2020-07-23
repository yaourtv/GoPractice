package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func onErrFail(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	onErrFail(err)
	for {
		message, err := bufio.NewReader(os.Stdin).ReadString('\n')
		onErrFail(err)
		if strings.Trim(message, "\r\n") == "exit" {
			return
		}
		fmt.Fprintf(conn, message)
		answer, err := bufio.NewReader(conn).ReadString('\n')
		onErrFail(err)
		fmt.Println(answer)
	}
}
