package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {
	defer conn.Close()
	Request(conn)
	Response(conn)
}

func Request(conn net.Conn) {

	i := 0
	scanner := bufio.NewScanner(conn) // scanning the connection
	for scanner.Scan() {              // return true idf scan is successful
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			m := strings.Fields(ln)[0] // takes the field of index 0 value
			u := strings.Fields(ln)[1]
			fmt.Println("**** METHOD", m)
			fmt.Println("****URL", u)
		}
		if ln == "" {
			break
		}
		i++
	}

}

func Response(conn net.Conn) {
	body := "<!DOCTYPE html><html><head><title>Title of the document</title></head><body>The content of the document......</body></html>"
	fmt.Fprint(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprint(conn, "Content-Length:\r\n", len(body))
	fmt.Fprint(conn, "Content Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
