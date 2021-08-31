package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloServer struct{}

func (p *HelloServer) Hello(request string, reply *string) error {
	*reply = "hushengs-hello:" + request
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloServer))
	listener, err := net.Listen("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("ListenTCP error", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error", err)
	}
	rpc.ServeConn(conn)
}
