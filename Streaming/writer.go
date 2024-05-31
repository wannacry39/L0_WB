package Streaming

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/stan.go"
)

func ReadMsg(sc stan.Conn, db *sql.DB, wg *sync.WaitGroup, m *sync.RWMutex) {
	m.RLock()
	sc.Subscribe("subj_str", func(m *stan.Msg) {
		res, err := db.Exec("insert into \"order\" (data) values ($1)", string(m.Data))

		if err != nil {
			log.Fatal(err)
		}

		defer wg.Done()
		fmt.Print(res)

	})
	m.RUnlock()
}
