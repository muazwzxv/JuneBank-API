package events

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/wagslane/go-rabbitmq"
)

type Payload struct {
	Data  []byte
	Event string
}

func (pl *Payload) ToBuffer() []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)

	err := enc.Encode(pl)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))

	return buf.Bytes()
}

func (p *Publisher) Publish(data *Payload, key []string) {
	// TODO: Message not published to queue, check the problem
	p.pub.Publish(
		data.ToBuffer(),
		key,
		rabbitmq.WithPublishOptionsContentType("application/json"),
		rabbitmq.WithPublishOptionsMandatory,
		rabbitmq.WithPublishOptionsPersistentDelivery,
		rabbitmq.WithPublishOptionsExchange("account-service"),
	)
}
