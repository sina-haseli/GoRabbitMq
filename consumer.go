package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("Consumer is running")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	defer func(conn *amqp.Connection) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer func(ch *amqp.Channel) {
		err := ch.Close()
		if err != nil {
			panic(err)
		}
	}(ch)

	msgs, err := ch.Consume(
		"hello",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			fmt.Printf("Received a message: %s", d.Body)
		}
	}()

	fmt.Println("connected and consuming")
	fmt.Println("Waiting for messages...")
	<-forever
}
