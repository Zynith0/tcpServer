package tcp

import (
	"fmt"
	"net"
)

func Server(port string) (net.Listener, error) {
	listener, err := net.Listen("tcp", port)
	fmt.Printf("Server listening on %v", listener.Addr())
	return listener, err
}

func Echo(conn net.Conn, message []byte) {
	conn.Write(message)
}
