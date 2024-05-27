package Stream

import (
	"time"

	"github.com/nats-io/stan.go"
)

func SendMsg(sc stan.Conn) {

	sc.Publish("subj_str", []byte("hello world"))
	time.Sleep(2 * time.Second)

}
