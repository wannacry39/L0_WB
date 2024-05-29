package Stream

import (
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func SendMsg(sc stan.Conn) {

	data, err := os.ReadFile("data/model.json")
	if err != nil {
		log.Fatal(err)
	}
	sc.Publish("subj_str", data)

}
