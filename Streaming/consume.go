package Stream

import (
	"fmt"

	"github.com/nats-io/stan.go"
)

func ReadMsg(sc stan.Conn) {
	sc.Subscribe("subj_str", func(m *stan.Msg) {
		fmt.Println(string(m.Data))
	})
}
