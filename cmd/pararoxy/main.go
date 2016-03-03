package main

import (
	"github.com/dinever/pararoxy"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3316")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go Pararoxy.Proxy(conn)
	}
}
