package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"redes/m/data"
	"redes/m/sorting"
	"strings"
)

type mapConn struct {
	conn    net.Conn
	request data.Request
}

func recvData(ctx context.Context, conn net.Conn,
	recv chan<- mapConn, deadConns chan<- net.Conn) {

	for {
		req := data.Request{}

		if err := req.Deserialize(conn); err != nil {
			break
		}

		recv <- mapConn{
			conn:    conn,
			request: req,
		}
	}

	select {
	case deadConns <- conn:
	case <-ctx.Done():
		return
	}
}

// connHandler -> Gerenciador de conexões do servidor
func connHandler(ctx context.Context, listener net.Listener,
	pool chan<- net.Conn) {
	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("Error on listener: %s\n", err.Error())
			return
		}

		log.Printf("Received a new connection from: %s\n", conn.RemoteAddr().String())

		select {
		case pool <- conn:
		case <-ctx.Done():
			log.Println("[Accept Handler] Cancelled by Context")
			return
		}
	}
}

// Server -> Implementação de um servidor tcp
func Server(ctx context.Context) {
	pool := make(chan net.Conn)
	recv := make(chan mapConn)
	deadConn := make(chan net.Conn)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()
	log.Println("Server running at localhost:8080")

	go connHandler(ctx, listener, pool)

	for {
		select {
		case conn := <-pool:
			go recvData(ctx, conn, recv, deadConn)
		case req := <-recv:
			res := data.Response{}

			switch strings.ToLower(req.request.SortMethod) {
			case "inserction":
				res.Error = nil
				res.Payload = sorting.InsertionSort(req.request.Payload...)
				break
			case "bubble":
				res.Error = nil
				res.Payload = sorting.BubbleSort(req.request.Payload...)
				break
			case "select":
				res.Error = nil
				res.Payload = sorting.SelectionSort(req.request.Payload...)
				break
			default:
				res.Error = fmt.Errorf("invalid sorting method")
				res.Payload = []int{}
			}

			serialized, _ := res.Serialize()
			req.conn.Write(serialized)
		case conn := <-deadConn:
			log.Printf("Connection %s closed...\n", conn.RemoteAddr().String())
		case <-ctx.Done():
			log.Println("[Server] Cancelled by Context")
			return
		}
	}
}
