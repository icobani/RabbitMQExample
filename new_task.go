/*
   © 2020 B1 Digital
   User    : ICI
   Name    : Ibrahim ÇOBANİ
   Date    : 10.06.2020  14:08
   Notes   :
*/
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
	conn, err := amqp.Dial("amqp://ici:aqswdefr+@rabbitmq.b1db.com:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"follow_notification_task_que", // name
		true,                           // durable => kuyruğun bir RabbitMQ düğümü yeniden başlatıldığında hayatta kalacağından emin olmalıyız.
		false,                          // delete when unused
		false,                          // exclusive
		false,                          // no-wait
		nil,                            // arguments
	)
	failOnError(err, "Failed to declare a queue")

	//body := bodyFrom(os.Args)
	body := ""
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}

//func bodyFrom(args []string) string {
//	var s string
//	if (len(args) < 2) || os.Args[1] == "" {
//		s = "hello"
//	} else {
//		s = strings.Join(args[1:], " ")
//	}
//	return s
//}
