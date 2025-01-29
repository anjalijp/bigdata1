package kafka

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bigdata1/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Review struct {
	ListingID    string `json:"listing_id"`
	ID           string `json:"id"`
	Date         string `json:"date"`
	ReviewerID   string `json:"reviewer_id"`
	ReviewerName string `json:"reviewer_name"`
	Comments     string `json:"comments"`
}

func Producer() {
	cfg := config.NewConfig()

	// Create a new Kafka producer
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBroker,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	// Open the CSV file
	file, err := os.Open(cfg.CSVFilePath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Skip the header row
	_, err = reader.Read()
	if err != nil {
		log.Fatalf("Failed to read CSV header: %v", err)
	}

	for {
		var chunk [][]string
		for i := 0; i < cfg.KafkaChunkSize; i++ {
			record, err := reader.Read()
			if err != nil {
				if err.Error() == "EOF" {
					break // End of file
				}
				log.Fatalf("Failed to read CSV record: %v", err)
			}
			chunk = append(chunk, record)
		}

		if len(chunk) == 0 {
			break // No more data to process
		}

		// Send the chunk to Kafka
		for _, record := range chunk {
			review := Review{
				ListingID:    string(record[0]),
				ID:           string(record[1]),
				Date:         string(record[2]),
				ReviewerID:   string(record[3]),
				ReviewerName: string(record[4]),
				Comments:     string(record[5]),
			}

			// Marshal the struct to JSON
			message, err := json.Marshal(review)
			if err != nil {
				log.Fatal(err)
			}

			hash := sha256.Sum256([]byte(message)) // Generate SHA-256 hash
			uniqueID := hex.EncodeToString(hash[:])
			err = producer.Produce(&kafka.Message{
				TopicPartition: kafka.TopicPartition{Topic: &cfg.KafkaTopic, Partition: kafka.PartitionAny},
				Key:            []byte(uniqueID),
				Value:          []byte(message),
			}, nil)

			if err != nil {
				log.Printf("Failed to produce message: %v\n", err)
			} else {
				fmt.Printf("Produced message: %s\n", message)
			}
		}
	}

	// Wait for all messages to be delivered
	producer.Flush(15 * 1000)
	fmt.Println("All messages produced successfully!")
}
