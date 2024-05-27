package main

import (
	// myServer "l0_project/Server"
	Stream "l0_project/Streaming"

	"github.com/nats-io/stan.go"
)

func main() {
	sc, _ := stan.Connect("test-cluster", "ID0")
	// myServer.Server()
	for {
		Stream.SendMsg(sc)
		Stream.ReadMsg(sc)

	}
}
