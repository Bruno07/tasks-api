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
