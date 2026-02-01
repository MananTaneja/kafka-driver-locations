package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	broker := "localhost:9092"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{broker},
		Topic:       "driver-location",
		Partition:   0,
		StartOffset: kafka.LastOffset,
		MinBytes:    10e3, // 10KB
		MaxBytes:    10e6, // 10MB
	})

	defer reader.Close()

	ctx := context.Background()

	fmt.Println("Consuming a single message from Kafka...")
	message, err := reader.ReadMessage(ctx)
	if err != nil {
		fmt.Printf("Failed to read message from Kafka: %v", err)
	}

	fmt.Printf("Message consumed: key=%s value=%s\n", string(message.Key), string(message.Value))

	fmt.Println("Kafka message consumption complete. Exiting.")
}
