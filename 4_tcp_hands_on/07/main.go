package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error")
	}

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatalln("Error")
			continue
		}

		go serve(conn)
	}
}

func serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	method, uri := "", ""

	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		if method == "" && uri == "" {
			spl := strings.Fields(ln)
			method = spl[0]
			uri = spl[1]
		}
	}
	defer conn.Close()

	body := `<h1>"We are in!"</h1>`
	io.WriteString(conn, "HTTP/1.1 200 OK \r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
	conn.Close()
}
