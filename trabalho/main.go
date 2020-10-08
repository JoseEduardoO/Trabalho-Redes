package main

import (
	"context"
	"redes/m/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server.Server(ctx)
}
