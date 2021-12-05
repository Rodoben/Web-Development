package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer lis.Close()

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
		}
		io.WriteString(conn, "Hello From TCP")
		fmt.Fprintln(conn, "How is your day")
		fmt.Fprintln(conn, " Jai Ho")
		conn.Close()

	}

}
