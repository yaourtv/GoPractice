package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

func onErrFail(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func processMessage(m string) string {
	if message, err := strconv.Atoi(m[:len(m)-2]); err == nil {
		message *= 2
		return strconv.Itoa(message)
	}
	return strings.ToUpper(m)
}

func main() {
	port := 8081
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	onErrFail(err)
	conn, err := ln.Accept()
	onErrFail(err)
	for {
		message, err := bufio.NewReader(conn).ReadString('\n')
		onErrFail(err)
		conn.Write([]byte(processMessage(message) + "\n"))
	}
}
