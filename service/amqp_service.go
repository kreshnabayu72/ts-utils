package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func logOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func getURI() string {
	hostname := os.Getenv("RABBITMQ_HOSTNAME")
	if hostname == "" {
		failOnError(errors.New("hostname can't be empty"), "Empty environment")
	}

	port := os.Getenv("RABBITMQ_PORT")
	if port == "" {
		failOnError(errors.New("port can't be empty"), "Empty environment")
	}

	return fmt.Sprintf("amqps://%s:%s/", hostname, port)
}

type QueueDeclareArgs struct {
	name       string
	durable    bool
	autoDelete bool
	exclusive  bool
	noWait     bool
	args       amqp.Table
}

type PublishArgs struct {
	exchange  string
	key       string
	mandatory bool
	immediate bool
	msg       amqp.Publishing
}

type ExchangeDeclareArgs struct {
	name       string
	kind       string
	durable    bool
	autoDelete bool
	internal   bool
	noWait     bool
	args       amqp.Table
}

type QueueBindArgs struct {
	name     string
	key      string
	exchange string
	noWait   bool
	args     amqp.Table
}

type ConsumeArgs struct {
	queue     string
	consumer  string
	autoAck   bool
	exclusive bool
	noLocal   bool
	noWait    bool
	args      amqp.Table
}

type Subscriber struct {
	Exchange   ExchangeDeclareArgs
	Queue      QueueDeclareArgs
	Bind       QueueBindArgs
	Consume    ConsumeArgs
	HandleFunc func(delivery <-chan amqp.Delivery)
}

type AMQPService interface {
	Publish(
		ctx context.Context,
		exchange ExchangeDeclareArgs,
		pub PublishArgs,
	) error
	Subscribe(
		exchange ExchangeDeclareArgs,
		queue QueueDeclareArgs,
		bind QueueBindArgs,
		consume ConsumeArgs,
		handleFunc func(delivery <-chan amqp.Delivery),
	) error
	SimplePublish(
		ctx context.Context,
		exchange string,
		key string,
		msg amqp.Publishing,
	) error
	SimpleSubscribe(
		exchange string,
		queue string,
		key string,
		consumer string,
		handleFunc func(delivery <-chan amqp.Delivery),
	) error
}

// AMQPServiceImpl is the implementation of the interface AMQPService,
// it abstracts the rabbitmq's amqp091-go package
type AMQPServiceImpl struct {
	URI               string
	Connection        *amqp.Connection
	ReconnectDelay    time.Duration
	ReconnectAttempts int
	VirtualHost       string
	Subscribers       map[uint64]Subscriber
	subIncrementId    uint64
}

// MakeAMQPService creates and initialize the connection of AMQPServiceImpl
func MakeAMQPService(
	reconnectDelay time.Duration,
	reconnectAttempts int,
	keepAlive bool,
) *AMQPServiceImpl {
	vhost := os.Getenv("VIRTUAL_HOST")
	if vhost == "" {
		failOnError(errors.New("vhost can't be empty"), "Empty environment")
	}

	amqpServiceImpl := AMQPServiceImpl{
		URI:               getURI(),
		ReconnectDelay:    reconnectDelay,
		ReconnectAttempts: reconnectAttempts,
		VirtualHost:       vhost,
		Subscribers:       make(map[uint64]Subscriber),
		subIncrementId:    0,
	}

	err := amqpServiceImpl.Connect()
	if err != nil {
		logOnError(
			err,
			fmt.Sprintf(
				"Could not connect to RabbitMQ after %d attempts",
				amqpServiceImpl.ReconnectAttempts,
			),
		)
	}

	if keepAlive {
		go amqpServiceImpl.keepAlive()
	}

	return &amqpServiceImpl
}

// Publish sends a message to an AMQP exchange
//
// It declares the exchange if it does not already exist
// and then publishes the message to that exchange with the specified routing key
func (s *AMQPServiceImpl) Publish(
	ctx context.Context,
	exchange ExchangeDeclareArgs,
	pub PublishArgs,
) error {
	ch, err := s.Connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err = ch.ExchangeDeclare(
		exchange.name,
		exchange.kind,
		exchange.durable,
		exchange.autoDelete,
		exchange.internal,
		exchange.noWait,
		exchange.args,
	); err != nil {
		return err
	}

	if err = ch.PublishWithContext(
		ctx,
		pub.exchange,
		pub.key,
		pub.mandatory,
		pub.immediate,
		pub.msg,
	); err != nil {
		return err
	}

	return nil
}

// Subscribe subscribes to messages from an AMQP exchange and queue using RabbitMQ
//
// sets up a subscription to receive messages from an AMQP exchange and queue.
// It declares the necessary components (exchange, queue, binding)
// and starts consuming messages with a specified handler function.
func (s *AMQPServiceImpl) Subscribe(
	exchange ExchangeDeclareArgs,
	queue QueueDeclareArgs,
	bind QueueBindArgs,
	consume ConsumeArgs,
	handleFunc func(delivery <-chan amqp.Delivery),
) error {
	ch, err := s.Connection.Channel()
	if err != nil {
		return err
	}

	if err = ch.ExchangeDeclare(
		exchange.name,
		exchange.kind,
		exchange.durable,
		exchange.autoDelete,
		exchange.internal,
		exchange.noWait,
		exchange.args,
	); err != nil {
		return err
	}

	if _, err = ch.QueueDeclare(
		queue.name,
		queue.durable,
		queue.autoDelete,
		queue.exclusive,
		queue.noWait,
		queue.args,
	); err != nil {
		return err
	}

	if err = ch.QueueBind(
		bind.name,
		bind.key,
		bind.exchange,
		bind.noWait,
		bind.args,
	); err != nil {
		return err
	}

	delivery, err := ch.Consume(
		consume.queue,
		consume.consumer,
		consume.autoAck,
		consume.exclusive,
		consume.noLocal,
		consume.noWait,
		consume.args,
	)
	if err != nil {
		return err
	}

	go handleFunc(delivery)

	// store subscriber for re-subscription on reconnection
	s.Subscribers[s.subIncrementId] = Subscriber{
		Exchange:   exchange,
		Queue:      queue,
		Bind:       bind,
		Consume:    consume,
		HandleFunc: handleFunc,
	}

	s.subIncrementId++

	return nil
}

// SimplePublish is the simplified version of Publish
func (s *AMQPServiceImpl) SimplePublish(
	ctx context.Context,
	exchange string,
	key string,
	msg amqp.Publishing,
) error {
	defaultExchange := ExchangeDeclareArgs{
		name:       exchange,
		kind:       "topic",
		durable:    true,
		autoDelete: false,
		internal:   false,
		noWait:     false,
		args:       nil,
	}

	defaultPublish := PublishArgs{
		exchange:  exchange,
		key:       key,
		mandatory: false,
		immediate: false,
		msg:       msg,
	}

	log.Printf("Publishing to: %s:%s", exchange, key)
	log.Printf("Message:\n%s", msg.Body)
	if err := s.Publish(ctx, defaultExchange, defaultPublish); err != nil {
		return err
	}

	return nil
}

// SimpleSubscribe is the simplified version of Subscribe
func (s *AMQPServiceImpl) SimpleSubscribe(
	exchange string,
	queue string,
	key string,
	consumer string,
	handleFunc func(delivery <-chan amqp.Delivery),
) error {
	defaultExchange := ExchangeDeclareArgs{
		name:       exchange,
		kind:       "topic",
		durable:    true,
		autoDelete: false,
		internal:   false,
		noWait:     false,
		args:       nil,
	}

	defaultQueue := QueueDeclareArgs{
		name:       queue,
		durable:    true,
		autoDelete: false,
		exclusive:  false,
		noWait:     false,
		args:       nil,
	}

	defaultBindQueue := QueueBindArgs{
		name:     queue,
		key:      key,
		exchange: exchange,
		noWait:   false,
		args:     nil,
	}

	defaultConsume := ConsumeArgs{
		queue:     queue,
		consumer:  consumer,
		autoAck:   true,
		exclusive: false,
		noLocal:   false,
		noWait:    false,
		args:      nil,
	}

	log.Printf("Subscribing to: %s:%s:%s", exchange, queue, key)
	if err := s.Subscribe(defaultExchange, defaultQueue, defaultBindQueue, defaultConsume, handleFunc); err != nil {
		return err
	}

	return nil
}

// Connect establishes a connection to a RabbitMQ server
//
// It implements a retry mechanism to handle connection failures
func (s *AMQPServiceImpl) Connect() error {
	var err error

	username := os.Getenv("RABBITMQ_USERNAME")
	password := os.Getenv("RABBITMQ_PASSWORD")

	auth := amqp.PlainAuth{
		Username: username,
		Password: password,
	}

	// When ReconnectAttempts <= 0, set ReconnectAttempts to max integer limit, basically means to attempt infinitely
	if s.ReconnectAttempts <= 0 {
		s.ReconnectAttempts = math.MaxInt
	}

	for i := 0; i < s.ReconnectAttempts; i++ {
		s.Connection, err = amqp.DialConfig(s.URI, amqp.Config{
			Vhost: s.VirtualHost,
			SASL: []amqp.Authentication{
				&auth,
			},
		})
		if err != nil {
			logOnError(
				err,
				fmt.Sprintf(
					"Failed to connect to RabbitMQ. Retrying in %v... (%d/%d)",
					s.ReconnectDelay,
					i+1,
					s.ReconnectAttempts,
				),
			)
			time.Sleep(s.ReconnectDelay)
			continue
		}
		log.Printf("Connected to RabbitMQ: %s", s.URI)

		return nil
	}

	logOnError(
		err,
		fmt.Sprintf("Could not connect to RabbitMQ after %d attempts", s.ReconnectAttempts),
	)

	return err
}

// keepAlive maintains a persistent connection to a RabbitMQ server
//
// It continuously monitors the connection for closures and attempts to reconnect if a closure is detected
func (s *AMQPServiceImpl) keepAlive() {
	notifyClose := make(chan *amqp.Error)
	s.Connection.NotifyClose(notifyClose)

	for {
		err := <-notifyClose
		if err != nil {
			log.Printf("Connection closed: %v. Reconnecting...", err)
			var reconnectErr error
			for {
				reconnectErr = s.Connect()
				if reconnectErr == nil {
					log.Println("Reconnected to RabbitMQ")
					// Re-assign notifyClose to the newly created connection
					notifyClose = make(chan *amqp.Error)
					s.Connection.NotifyClose(notifyClose)

					// Attempts to re-subscribe stored subscriber(s)
					s.Resubscribe()
					break
				}
				log.Printf(
					"Reconnection failed: %v. Retrying in %v...",
					reconnectErr,
					s.ReconnectDelay,
				)
				time.Sleep(s.ReconnectDelay)
			}
		}
	}
}

func (s *AMQPServiceImpl) Resubscribe() {
	for id, subscriber := range s.Subscribers {
		go func() {
			err := s.Subscribe(
				subscriber.Exchange,
				subscriber.Queue,
				subscriber.Bind,
				subscriber.Consume,
				subscriber.HandleFunc,
			)
			if err != nil {
				logOnError(err, fmt.Sprintf("Failed to subscribe: %s", err))
			}
		}()

		// Delete previous subscriber because it's already been rejuvenated with another id
		delete(s.Subscribers, id)
	}
}
