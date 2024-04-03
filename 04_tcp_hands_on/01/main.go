package main

import (
	"io"
	"log"
	"net"
)

func main() {
	conn, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error")
	}

	defer conn.Close()

	for {
		req, err := conn.Accept()
		if err != nil {
			log.Fatalln("Error")
		}
		io.WriteString(req, "I see you connected")
		req.Close()
	}
}
