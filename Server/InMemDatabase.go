package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {

	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	defer lis.Close()
	for {
		conn, err := lis.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}

}

func handle(conn net.Conn) {

	defer conn.Close()
	io.WriteString(conn, "\r\nIN-MEMORY DATABASE\r\n\r\n"+
		"USE:\r\n"+
		"\tSET key value \r\n"+
		"\tGET key \r\n"+
		"\tDEL key \r\n\r\n"+
		"EXAMPLE:\r\n"+
		"\tSET fav chocolate \r\n"+
		"\tGET fav \r\n\r\n\r\n")

	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s\n", v)

		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED_VALUE")
				continue
			}

			k := fs[1]
			v := fs[2]
			data[k] = v

		case "DEL":
			k := fs[1]
			delete(data, k)

		default:
			fmt.Fprintln(conn, "Invalid Command")
		}
	}

}
