package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/streadway/amqp"
)

//this is producer
func main() {
	//start connection
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

	//declare queue first
	q2, err := channel.QueueDeclare(
		"hey", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)

	count := 0
	num := 0
	for {
		num = num % 2
		fmt.Println(num)
		if num == 0 {
			//demo to send message to queue "hello"
			err = channel.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(strconv.Itoa(count) + "Hello World!"),
				})
		} else {
			err = channel.Publish(
				"",      // exchange
				q2.Name, // routing key
				false,   // mandatory
				false,   // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(strconv.Itoa(count) + "hey World!"),
				})
		}
		num = rand.Intn(100000)
		fmt.Println(num)
		fmt.Println("Message sent.")
		time.Sleep(time.Duration(num))
		count++
	}
}
