package message_broker

import (
	"github.com/streadway/amqp"
	"reservation-api/pkg/applogger"
)

type MessageBrokerManager interface {

	// PublishMessage publish message in given queue name
	PublishMessage(qName string, payload []byte) error
	// Consume consumes message in given queue name
	Consume(qName string, fn func(payload interface{})) error
}

// RabbitMQManager implements MessageBrokerManager interface.
type RabbitMQManager struct {
	Connection *amqp.Connection
	Logger     applogger.Logger
}

// New returns new instance of RabbitMQManager
// This function has two parameters called logger and url.
// If it fails to dial the given url, it will panic.
func New(url string, logger applogger.Logger) *RabbitMQManager {
	con, err := amqp.Dial(url)
	//defer con.Close()

	if err != nil {
		logger.LogError(err.Error())
		panic(err.Error())
	}

	return &RabbitMQManager{
		Connection: con,
		Logger:     logger,
	}
}

// PublishMessage publishes message to given queue name.
func (m *RabbitMQManager) PublishMessage(qName string, payload []byte) error {

	ch, err := getChannel(qName, m.Connection)
	defer ch.Close()

	if err != nil {
		return err
	}

	return ch.Publish("", qName, false, false, amqp.Publishing{
		Headers:     nil,
		ContentType: "text/plain",
		Body:        payload,
	})
}

// Consume consumes message firm given queue name
// and runs given function after consume message.
func (m RabbitMQManager) Consume(qName string, fn func(payload interface{})) error {

	ch, err := getChannel(qName, m.Connection)
	defer ch.Close()

	if err != nil {
		return err
	}

	delivery, err := ch.Consume(qName, "", true, false, false, false, nil)

	if err != nil {
		m.Logger.LogError(err.Error())
		return err
	}

	go func() {
		for msg := range delivery {
			fn(msg)
		}
	}()

	return nil
}

/*=========================  private functions =================================*/

func getChannel(qName string, con *amqp.Connection) (*amqp.Channel, error) {

	ch, err := con.Channel()
	if err != nil {
		return nil, err
	}
	ch.QueueDeclare(
		qName,
		false, false, false, false, nil,
	)
	return ch, nil
}
