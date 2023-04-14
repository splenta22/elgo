package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ravsii/elgo"
	"github.com/ravsii/elgo/examples/queue/player"
)

func main() {
	pool := elgo.NewPool(elgo.WithRetry(1*time.Second), elgo.WithIncreaseInterval(0.03))

	for i := 0; i < 1000; i++ {
		pool.AddPlayer(&player.Player{Name: fmt.Sprint(i), EloRating: rand.Float64()})
	}

	fmt.Println("pool size", pool.Size())

	go pool.Run()

	for i := 0; i < 500; i++ {
		match := <-pool.Matches()
		p1 := match.Player1.(*player.Player)
		p2 := match.Player2.(*player.Player)
		fmt.Println(i+1, "match", p1, p2)
	}

	fmt.Println(pool.Close())
}
