package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"cloud.google.com/go/pubsub"
)

func main() {
	ctx := context.Background()

	// Configura el cliente de Pub/Sub.
	client, err := pubsub.NewClient(ctx, "tss-1s2024")
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}

	// Define el nombre del tema y la suscripción.
	topicName := "my-topic"
	subscriptionName := "my-sub"

	// Publica un mensaje en el tema.
	if err := publishMessage(ctx, client, topicName, "Hello, Pub/Sub!"); err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	// Espera un momento para que Pub/Sub procese el mensaje.
	time.Sleep(1 * time.Second)

	// Recibe mensajes desde la suscripción.
	if err := receiveMessages(ctx, client, subscriptionName); err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}
}

func publishMessage(ctx context.Context, client *pubsub.Client, topicName, message string) error {
	// Obtiene el tema.
	topic := client.Topic(topicName)

	// Publica el mensaje.
	result := topic.Publish(ctx, &pubsub.Message{
		Data: []byte(message),
	})

	// Espera a que la publicación se complete.
	_, err := result.Get(ctx)
	return err
}

func receiveMessages(ctx context.Context, client *pubsub.Client, subscriptionName string) error {
	// Obtiene la suscripción.
	sub := client.Subscription(subscriptionName)

	// Configura el handler para procesar los mensajes recibidos.
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	err := sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		// Procesa el mensaje recibido.
		fmt.Printf("Received message: %s\n", string(msg.Data))
		msg.Ack() // Acknowledge the message.
	})
	return err
}
