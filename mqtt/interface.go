package mqtt

type MQTTMessageHandler func(topic string, retained bool, payload []byte)

// MQTTService defines the interface for interacting with an MQTT broker.
type MQTTService interface {
	// Connect establishes a connection to the MQTT broker.
	Connect() error

	// Subscribe subscribes to a given topic and registers a handler function
	// to process incoming messages on that topic.
	Subscribe(topic string, handler MQTTMessageHandler) error

	// Publish sends a message to a specified topic.
	Publish(topic string, message []byte) error

	// Unsubscribe removes the subscription to a given topic.
	Unsubscribe(topic string) error
}
