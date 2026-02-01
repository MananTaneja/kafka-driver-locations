package main

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "localhost:9092"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{brokerAddress},
		Topic:     "driver-location",
		Balancer:  &kafka.LeastBytes{},
		BatchSize: 1,
	})

	defer writer.Close()

	message := kafka.Message{
		Key:   []byte("driver-key"),
		Value: []byte("12"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	fmt.Println("Sending message to kafka")
	err := writer.WriteMessages(ctx, message)

	if err != nil {
		fmt.Printf("error occured: %v", err)
	}

	fmt.Println("Message sent successfully.")
	fmt.Println("Kafka producer has completed sending the message.")
}
