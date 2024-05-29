package main

import (
	// myServer "l0_project/Server"
	"fmt"
	Stream "l0_project/Streaming"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "ID0")
	// myServer.Server()
	ch := make(chan []byte)

	go Stream.SendMsg(sc)
	go Stream.ReadMsg(sc, ch)
	fmt.Print(string(<-ch))

}
