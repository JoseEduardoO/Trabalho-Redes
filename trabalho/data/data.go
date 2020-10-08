package data

import (
	"encoding/json"
	"net"
)

// Manip -> Interface de manipulação de dados
type Manip interface {
	Serialize() ([]byte, error)
	Deserialize(net.Conn) error
}

// Request -> Estrutura para uma requisição.
type Request struct {
	SortMethod string
	Payload    []int
}

// Response -> Estrutura para uma resposta.
type Response struct {
	Payload []int
	Error   error
}

// NewRequest -> Criação de uma nova requisição.
func NewRequest(sortMethod string, values ...int) *Request {
	return &Request{
		SortMethod: sortMethod,
		Payload:    values,
	}
}

// NewResponse -> Criação de uma resposta
func NewResponse(err error, payload ...int) *Response {
	return &Response{
		Payload: payload,
		Error:   err,
	}
}

// Serialize -> Implementação da função de serialização por parte do cliente.
func (req *Request) Serialize() ([]byte, error) {
	return json.Marshal(&req)
}

// Deserialize -> Implementação da função de deserialização por parte do cliente.
func (req *Request) Deserialize(conn net.Conn) error {
	decoder := json.NewDecoder(conn)
	return decoder.Decode(&req)
}

// Serialize -> Implementação da função de deserialização por parte do servidor.
func (res *Response) Serialize() ([]byte, error) {
	return json.Marshal(&res)
}

// Deserialize -> Implementação da função de deserialização por parte do servidor.
func (res *Response) Deserialize(conn net.Conn) error {
	decoder := json.NewDecoder(conn)
	return decoder.Decode(&res)
}
