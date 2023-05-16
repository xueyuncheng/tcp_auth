package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:9000")
	if err != nil {
		log.Fatalln(err)
	}

	if _, err := conn.Write([]byte("hello")); err != nil {
		log.Fatalln(err)
	}

	buf := make([]byte, 1000)
	if _, err := conn.Read(buf); err != nil {
		log.Fatalln(err)
	}

	log.Println(string(buf))
}
