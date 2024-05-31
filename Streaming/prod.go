package Streaming

import (
	"log"
	"os"
	"sync"

	"github.com/nats-io/stan.go"
)

func SendMsg(sc stan.Conn, wg *sync.WaitGroup, m *sync.RWMutex) {

	data, err := os.ReadFile("data/model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer wg.Done()

	m.Lock()
	sc.Publish("subj_str", data)
	m.Unlock()

}
