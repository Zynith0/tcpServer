package tcp

import (
	"fmt"
	"net"
	"sync"
)

type Server struct {
	listener net.Listener
	clients map[net.Addr]net.Conn
	mutex sync.Mutex
}

// var (
// 	clients = make(map[string]*Client)
// 	mutex sync.Mutex
// )


func (s *Server) CreateServer(port string) (*Server, error) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}
	return &Server{
		listener: listener,
		clients: make(map[net.Addr]net.Conn),
	}, nil
}

func (s *Server) Start(handler func(conn net.Conn)) error {
	fmt.Printf("Server listening on %v", s.listener.Addr())

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			return err
		}
		defer conn.Close()

		handler(conn)
	}
}

func Echo(conn net.Conn, message []byte) {
	conn.Write(message)
}

// func HandleConnection(conn net.Conn) {
// 	c := &Client{
// 		ID: conn.RemoteAddr().String(),
// 		Conn: conn,
// 	}
//
// 	fmt.Printf("Client %v has connected", c.ID)
//
// 	mutex.Lock()
// 	clients[c.ID] = c
// 	mutex.Unlock()
//
// 	// defer func() {
// 	// 	mutex.Lock()
// 	// 	delete(clients, c.ID)
// 	// 	mutex.Unlock()
// 	// 	conn.Close()
// 	// }()
// }

func (s *Server) Broadcast(message []byte) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	for _, client := range s.clients {
		_, err := client.Write(message)
		if err != nil {
			fmt.Println("skill issue", err)
		}
	}
}
