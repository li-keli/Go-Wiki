package main

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@127.0.0.1:5672/test")
	failOnError(err, "conn error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "channel open error ")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "exchange declare error")

	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "queue declare error")

	err = ch.QueueBind(
		q.Name,
		"",
		"logs",
		false,
		nil)
	failOnError(err, "queue bind error")

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "consumer register error")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" info -> %s", d.Body)
		}
	}()

	log.Printf("Waiting ...")
	<-forever
}
