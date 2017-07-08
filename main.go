package main

import (
	"log"

	nats "github.com/nats-io/go-nats"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	recvCh := make(chan string)
	ec.BindRecvChan("hello", recvCh)

	// Receive via Go channels
	for {
		who := <-recvCh
		log.Println("received ", who)
	}
}
