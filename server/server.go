package main

import (
	"bytes"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		do(conn)
	}
}

func do(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 5)
	if _, err := conn.Read(buf); err != nil {
		log.Println(err)
		return
	}

	if !bytes.Equal(buf, []byte("hello")) {
		conn.Write([]byte("wrong password"))
		return
	}

	if _, err := conn.Write([]byte("right password")); err != nil {
		log.Println(err)
		return
	}
}
