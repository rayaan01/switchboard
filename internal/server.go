package internal

import (
	"fmt"
	"net"
)

type Server struct {
	address  string
	listener net.Listener
}

type ServerHandler func(conn net.Conn, server *Server)

func CreateServer(host string, port uint16) (*Server, error) {
	address := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return nil, err
	}
	fmt.Printf("server listening on %s \n", address)
	serverInstance := Server{address, listener}
	return &serverInstance, nil
}

func (s *Server) AcceptConnections(handler ServerHandler) {
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			conn.Close()
			fmt.Println("Could not accept connection", err)
			continue
		}
		go handler(conn, s)
	}
}
