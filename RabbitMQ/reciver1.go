package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {
	//open connection
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//open channel
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer channel.Close()

	//declare queue first
	q, err := channel.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)

	//declare want to receive message from queue "hello"
	msgs, err := channel.Consume(
		q.Name, // queue (q.Name is "hello")
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		fmt.Println(err)
	}

	//using go routine to wait msgs to receive
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	//blocking program in here, because we still want to receive value from msgs
	cforever := make(chan bool)
	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-cforever
}
