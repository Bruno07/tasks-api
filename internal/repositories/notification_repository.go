package repositories

import amqp "github.com/rabbitmq/amqp091-go"

type INotificationRepository interface {
	Notify(payload []byte, exchange string, routingKey string) error
	Consumer(queue string, consumer string) (<-chan amqp.Delivery, error)
}

type NotificationRepository struct {
	queue *amqp.Channel
}

type MockNotificationRepository struct {
	MockNotify   func(payload []byte, exchange string, routingKey string) error
	MockConsumer func(queue string, consumer string) (<-chan amqp.Delivery, error)
}

func NewNotificationRepository(queue *amqp.Channel) *NotificationRepository {
	return &NotificationRepository{
		queue: queue,
	}
}

// Send message to queue
func (nr *NotificationRepository) Notify(payload []byte, exchange string, routingKey string) error {

	err := nr.queue.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(payload),
		},
	)

	return err

}

// Consumes message from queue
func (nr *NotificationRepository) Consumer(queue string, consumer string) (<-chan amqp.Delivery, error) {

	q, err := nr.queue.QueueDeclare(
		queue,
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return nil, err
	}

	msgs, err := nr.queue.Consume(
		q.Name,
		consumer,
		true,
		false,
		false,
		false,
		nil,
	)

	return msgs, err

}

// Send message to queue (Mock)
func (nr *MockNotificationRepository) Notify(payload []byte, exchange string, routingKey string) error {
	return nr.MockNotify(payload, exchange, routingKey)
}

// Consumes message from queue (Mock)
func (nr *MockNotificationRepository) Consumer(queue string, consumer string) (<-chan amqp.Delivery, error) {
	return nr.MockConsumer(queue, consumer)
}
