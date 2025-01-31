package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bigdata1/config"
	"github.com/bigdata1/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Consumer() {

	cfg := config.NewConfig()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     cfg.KafkaBroker,
		"broker.address.family": "v4",
		"group.id":              cfg.KafkaConsumerGroup,
		"session.timeout.ms":    6000,
		"auto.offset.reset":     "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics([]string{cfg.KafkaTopic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic: %v", err)
	}

	fmt.Println("Consumer is running...")
	for {
		msg, err := c.ReadMessage(-1)
		if err != nil {
			log.Printf("Consumer error: %v\n", err)
			continue
		}

		var review models.Review
		err = json.Unmarshal(msg.Value, &review)
		if err != nil {
			log.Printf("Failed to unmarshal message: %v\n", err)
			continue
		}

		fmt.Printf("Consumed message: %+v\n", review)
	}
}
