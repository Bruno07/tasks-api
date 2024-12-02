package queue

import (
	"sync"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	queueInstance *amqp.Channel
	once          sync.Once
)

// Connect to rabbitmq
func Connect(dns string) {

	once.Do(func() {

		conn, err := amqp.Dial(dns)
		if err != nil {
			panic(err.Error())
		}
	
		ch, err := conn.Channel()
		if err != nil {
			panic(err.Error())
		}

		queueInstance = ch

	})

}

func GetInstanceQueue() *amqp.Channel {
	return queueInstance
}

// Send message to queue
func Notify(payload []byte, exchange string, routingKey string, ch *amqp.Channel) {
	err := ch.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	if err != nil {
		panic(err.Error())
	}
}

// Consumes message from queue
func Consumer(queue string, consumer string, ch *amqp.Channel, in chan []byte) {
	q, err := ch.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		panic(err.Error())
	}

	go func() {
		for m := range msgs {
			in <- []byte(m.Body)
		}
		close(in)
	}()
}
