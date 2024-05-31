package main

import (
	// myServer "l0_project/Server"
	"database/sql"
	"l0_project/Streaming"
	"log"
	"sync"

	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
)

type sync_params struct {
	wg sync.WaitGroup
	m  sync.RWMutex
}

func main() {
	sc, _ := stan.Connect("test-cluster", "ID0")
	s := sync_params{}
	// myServer.Server()
	connstr := "user=user_l0 password=qwerty host=localhost dbname=l0 sslmode=disable"
	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	s.wg.Add(2)
	go Streaming.SendMsg(sc, &s.wg, &s.m)
	go Streaming.ReadMsg(sc, db, &s.wg, &s.m)
	s.wg.Wait()
}
