package main

import (
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Установите соединение с сервером NATS Streaming
	sc, err := stan.Connect("test-cluster", "consumer", stan.NatsURL("nats://localhost:4222"))
	if err != nil {
		log.Fatal(err)
	}
	defer sc.Close()

	// Функция обратного вызова для обработки входящих сообщений
	messageHandler := func(msg *stan.Msg) {
		log.Printf("Received a message: %s\n", msg.Data)
	}

	// Подписаться на канал "foo" с использованием функции обратного вызова
	sub, err := sc.Subscribe("foo", messageHandler)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Ждем сигнала для выхода
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	log.Println("Consumer is shutting down")
}
