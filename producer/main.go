package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	n := flag.Int("n", 5, "number of messages to send")
	flag.Parse()

	brokerAddress := "localhost:9092"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:   []string{brokerAddress},
		Topic:     "driver-location",
		Balancer:  &kafka.LeastBytes{},
		BatchSize: 1,
	})

	defer writer.Close()

	for i := 0; i < *n; i++ {
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

	fmt.Printf("Sent %d messages. Done.\n", *n)
}
