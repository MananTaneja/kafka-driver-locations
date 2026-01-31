package main

import (
	"github.com/segmentio/kafka-go"
)

func main() {
	brokerAddress := "localhost:9092"

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
	})

	defer writer.Close()
}
