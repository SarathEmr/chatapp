package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {
	log.Print("welcome to kafkaconsumer !")

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"kafka:9092"}, // Use the service name from docker-compose
		Topic:   "new_employee",
		// GroupID: "employee-consumer-group", // Consumer group for offset management
	})
	defer kafkaReader.Close()

	for {
		time.Sleep(5 * time.Second)
		msg, err := kafkaReader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("could not read kafka message: %v", err)
			continue
		}
		log.Printf("kafka received: key %s, value %s", string(msg.Key), string(msg.Value))
	}

}
