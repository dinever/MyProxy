package Pararoxy

import (
	"io"
	"log"
	"net"
)

func Proxy(conn net.Conn) {
	server, err := net.Dial("tcp", "127.0.0.1:3306")
	if err != nil {
		log.Println("Failed to dail the server.")
		log.Println(err)
		conn.Close()
		return
	}
	go receive(server, conn)
	go send(server, conn)
	// Sleep forever.
	for {
		select {}
	}
	server.Close()
	conn.Close()
}

func receive(server, client net.Conn) {
	buffer := make([]byte, maxPacketLength)
	for {
		n, err := server.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}
		_, err = client.Write(buffer[:n])
		if err != nil && err != io.EOF {
			return
		}
	}
}

func send(server, client net.Conn) {
	buffer := make([]byte, maxPacketLength)
	for {
		n, err := client.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}
		switch buffer[4] {
		case comQuery:
			println(string(buffer[5:n]))
		}
		_, err = server.Write(buffer[:n])
		if err != nil && err != io.EOF {
			return
		}
	}
}
