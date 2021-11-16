package tcp

import (
	"dozn/transaction-server/handlers"
	"fmt"
	"net"
)

func TcpServer() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to Listen : ", err)
	}

	fmt.Println("Server start !!")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to Accept : ", err)
			continue
		}
		fmt.Println("Client connect")

		go handlers.Handler(conn)
	}
}
