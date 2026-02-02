# Kafka Driver Locations

A simple Kafka exercise where a producer streams driver location events and a consumer processes them.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.18+)
- [Docker](https://docs.docker.com/get-docker/) and [Docker Compose](https://docs.docker.com/compose/install/)

## Getting Started

### 1. Start Kafka

Use Docker Compose to start a local Kafka broker:

```bash
docker-compose up -d
```

### 2. Run the Consumer

In a new terminal, start the consumer to wait for messages:

```bash
cd consumer
go run main.go
```

### 3. Run the Producer

In another terminal, run the producer to send a driver location event:

```bash
cd producer
go run main.go
```

## Project Structure

- `producer/`: Sends driver location events to the `driver-location` Kafka topic.
- `consumer/`: Reads events from the `driver-location` topic and logs them.
- `docker-compose.yml`: Configures a single-node Kafka cluster for local development.
