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
	failOnError(err, "connect error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "channel open error")
	defer ch.Close()

	body := "I want to go to WC!!!"
	err = ch.Publish(
		"logs",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "publish error")

	log.Printf(body)
}
