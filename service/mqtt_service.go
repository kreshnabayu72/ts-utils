package service

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type MQTTMessageHandler func(topic string, retained bool, payload []byte)

// MQTTServicer defines the interface for interacting with an MQTT broker.
type MQTTServicer interface {
	Init()
	// Connect establishes a connection to the MQTT broker.
	Connect() error
	// Subscribe subscribes to a given topic and registers a handler function
	// to process incoming messages on that topic.
	Subscribe(topic string, handler MQTTMessageHandler) error
	// Publish sends a message to a specified topic.
	Publish(topic string, message []byte) error
	// Unsubscribe removes the subscription to a given topic.
	Unsubscribe(topic string) error
	Close()
}

// MQTTService implements MQTTServicer
//
// Implementation of the MQTTServicer interface.
type MQTTService struct {
	client mqtt.Client // The MQTT client instance.
}

// MakeMQTTService creates and initializes an MQTTService instance.
// It connects to the MQTT broker using the provided credentials.
//
// Parameters:
// - broker: The MQTT broker URL.
// - clientID: The client ID to identify the connection.
// - username: The username for authentication.
// - password: The password for authentication.
//
// Returns:
// - A pointer to an initialized MQTTService instance.
// - An error if the connection to the broker fails.
func MakeMQTTService(broker, clientID, username, password string) *MQTTService {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	return &MQTTService{client: client}
}

func (m *MQTTService) Init() {
	// Pass
}

// Connect establishes a connection to the MQTT broker.
//
// Returns:
// - An error if the connection to the broker fails.
func (m *MQTTService) Connect() error {
	if token := m.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Subscribe subscribes to a topic and registers a handler function to process
// incoming messages.
//
// Parameters:
// - topic: The topic to subscribe to.
// - handler: A function to handle incoming messages.
//
// Returns:
// - An error if the subscription fails.
func (m *MQTTService) Subscribe(topic string, handler MQTTMessageHandler) error {
	if token := m.client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		handler(msg.Topic(), msg.Retained(), msg.Payload())
	}); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

// Publish sends a message to a specified topic.
//
// Parameters:
// - topic: The topic to publish the message to.
// - message: The message payload to send.
//
// Returns:
// - An error if the publish operation fails.
func (m *MQTTService) Publish(topic string, message []byte) error {
	token := m.client.Publish(topic, 0, false, message)
	token.Wait()
	return token.Error()
}

// Unsubscribe removes the subscription to a given topic.
//
// Parameters:
// - topic: The topic to unsubscribe from.
//
// Returns:
// - An error if the unsubscription fails.
func (m *MQTTService) Unsubscribe(topic string) error {
	if token := m.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (m *MQTTService) Close() {
	// Pass
}
