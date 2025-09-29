package tcp

import (
	"fmt"
	"net"
	"sync"
)

type Client struct {
	ID string
	Conn net.Conn
	Username string
}

var (
	clients = make(map[string]*Client)
	mutex sync.Mutex
)


func Server(port string) (net.Listener, error) {
	listener, err := net.Listen("tcp", port)
	fmt.Printf("Server listening on %v", listener.Addr())
	return listener, err
}

func Echo(conn net.Conn, message []byte) {
	conn.Write(message)
}

func HandleConnection(conn net.Conn) {
	c := &Client{
		ID: conn.RemoteAddr().String(),
		Conn: conn,
	}

	mutex.Lock()
	clients[c.ID] = c
	mutex.Unlock()

	// defer func() {
	// 	mutex.Lock()
	// 	delete(clients, c.ID)
	// 	mutex.Unlock()
	// 	conn.Close()
	// }()
}

func Broadcast(message []byte) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, client := range clients {
		_, err := client.Conn.Write(message)
		if err != nil {
			fmt.Println("skill issue", err)
		}
	}
}
