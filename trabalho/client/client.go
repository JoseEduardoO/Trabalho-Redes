package main

import (
	"context"
	"log"
	"net"
	"redes/m/data"
)

func sendMsg(ctx context.Context, dial net.Conn, payload []byte,
	recv chan<- data.Response) {
	dial.Write(payload)
	res := data.Response{}

	for {
		err := res.Deserialize(dial)

		if err != nil {
			return
		}

		select {
		case recv <- res:
		case <-ctx.Done():
			return
		}
	}
}

// client -> Implementação de um cliente tcp
func client(ctx context.Context, method string, vetor ...int) {
	recv := make(chan data.Response)
	req := data.NewRequest(method, vetor...)

	dial, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer dial.Close()

	serialized, err := req.Serialize()

	if err != nil {
		log.Fatal(err)
	}

	go sendMsg(ctx, dial, serialized, recv)

	select {
	case res := <-recv:
		if res.Error != nil {
			log.Fatal(res.Error)
		}

		log.Printf("(%s) Response from %s: %v\n", method,
			dial.RemoteAddr().String(), res.Payload)
		return
	case <-ctx.Done():
		log.Println("[Client] Cancelled by Context")
		return
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	//	client(ctx, "bubble", -10, 2, 0, 1, 5, 3, -2)
	client(ctx, "select", -10, 2, 0, 1, 5, 3, -2)
	//	client(ctx, "inserction", -10, 2, 0, 1, 5, 3, -2)

}
