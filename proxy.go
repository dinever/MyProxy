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
	go forward(server, conn)
	go forward(conn, server)
	// Sleep forever.
	for {
		select {}
	}
	server.Close()
	conn.Close()
}

func forward(src, dest net.Conn) {
	buffer := make([]byte, 1024)
	for {
		n, err := src.Read(buffer)
		if err != nil && err != io.EOF {
			return
		}
		_, err = dest.Write(buffer[0:n])
		if err != nil && err != io.EOF {
			return
		}
	}
}
