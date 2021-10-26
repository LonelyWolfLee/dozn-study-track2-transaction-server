package handlers

import (
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn) {
	recv := make([]byte, 4096)

	for {
		n, err := conn.Read(recv)
		if err != nil {
			if err == io.EOF {
				fmt.Println("connection is closed from client : ", conn.RemoteAddr().String())
			}
			fmt.Println("Failed to receive data : ", err)
			break
		}

		if n > 0 {
			fmt.Println(string(recv[:n]))
			conn.Write(recv[:n])
		}
	}
}
