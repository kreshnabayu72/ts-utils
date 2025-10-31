package mqtt

import mqtt "github.com/eclipse/paho.mqtt.golang"

// MQTTServiceImpl is the implementation of the MQTTService interface.
type MQTTServiceImpl struct {
	client mqtt.Client // The MQTT client instance.
}

// MakeMQTTService creates and initializes an MQTTServiceImpl instance.
// It connects to the MQTT broker using the provided credentials.
//
// Parameters:
// - broker: The MQTT broker URL.
// - clientID: The client ID to identify the connection.
// - username: The username for authentication.
// - password: The password for authentication.
//
// Returns:
// - A pointer to an initialized MQTTServiceImpl instance.
// - An error if the connection to the broker fails.
func MakeMQTTService(broker, clientID, username, password string) *MQTTServiceImpl {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	return &MQTTServiceImpl{client: client}
}

// Connect establishes a connection to the MQTT broker.
//
// Returns:
// - An error if the connection to the broker fails.
func (m *MQTTServiceImpl) Connect() error {
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
func (m *MQTTServiceImpl) Subscribe(topic string, handler MQTTMessageHandler) error {
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
func (m *MQTTServiceImpl) Publish(topic string, message []byte) error {
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
func (m *MQTTServiceImpl) Unsubscribe(topic string) error {
	if token := m.client.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
