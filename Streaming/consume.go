package Stream

import (
	"github.com/nats-io/stan.go"
)

func ReadMsg(sc stan.Conn, ch chan []byte) {
	sc.Subscribe("subj_str", func(m *stan.Msg) {
		ch <- m.Data
	})

}
