package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	brokerAddress := "localhost:9092"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{brokerAddress},
		Topic:     "driver-location",
		Balancer:  &kafka.LeastBytes{},
		BatchSize: 1,
	})

	defer writer.Close()

	for {
		randomValue := fmt.Sprintf("%d", rand.Intn(100))
		message := kafka.Message{
			Key:   []byte("driver-key"),
			Value: []byte(randomValue),
		}

		fmt.Printf("Sending message: %s\n", randomValue)
		err := writer.WriteMessages(context.Background(), message)

		if err != nil {
			fmt.Printf("error occured: %v\n", err)
		} else {
			fmt.Println("Message sent successfully.")
		}

		time.Sleep(2 * time.Second)
	}
}
