package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		con, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(con)
	}

}

func handle(con net.Conn) {
	scanner := bufio.NewScanner(con)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer con.Close()

	fmt.Println("We never get here.")
}
