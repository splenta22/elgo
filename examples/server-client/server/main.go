package main

import (
	"log"

	"github.com/ravsii/elgo"
	"github.com/ravsii/elgo/examples/player"
	"github.com/ravsii/elgo/socket"
)

func main() {
	pool := elgo.NewPool()
	pool.AddPlayer(
		player.New("Example 1", 0),
		player.New("Example 2", 0))

	defer pool.Close()

	go pool.Run()

	server := socket.NewServer(":3000", pool)
	log.Fatal(server.Listen())
}
