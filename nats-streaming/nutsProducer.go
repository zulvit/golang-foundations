package main

import (
	stan "github.com/nats-io/stan.go"
	"log"
)

func main() {
	sc, err := stan.Connect("test-cluster", "producer", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Опубликуйте сообщение в канал "foo"
	err = sc.Publish("foo", []byte("Hello, NATS Streaming from producer!"))
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Message published")
}
