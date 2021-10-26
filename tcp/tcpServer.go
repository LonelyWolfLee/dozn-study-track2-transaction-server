package tcp

import (
	"dozn/transaction-server/handlers"
	"fmt"
	"net"
)

func tcpServer() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Failed to Listen : ", err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to Accept : ", err)
			continue
		}

		handlers.handler(conn)
	}
}
