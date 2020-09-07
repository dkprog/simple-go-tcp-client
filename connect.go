package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "dkprog.com:80")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(conn, "HEAD / HTTP/1.1\r\nHost: dkprog.com\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(status)
}
