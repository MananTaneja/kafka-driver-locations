package main

import (
	"github.com/segmentio/kafka-go"
)

func main() {
	broker := "localhost:9092"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   "driver-location",
	})

	defer reader.Close()
}
