package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	println("Go ran successfully")

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
	fmt.Println("Connected to RabbitMQ")

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
	fmt.Println("Channel created")

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	fmt.Println(q)

	if err != nil {
		panic(err)
	}

	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("KIR TOSH!"),
		})

	if err != nil {
		panic(err)
	}

	fmt.Println("Message sent")
}
