package producer

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/bigdata1/config"
	"github.com/bigdata1/models"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Producer() {
	cfg := config.NewConfig()

	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.KafkaBroker,
	})
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	file, err := os.Open(cfg.CSVFilePath)
	if err != nil {
		log.Fatalf("Failed to open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

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
					break
				}
				log.Fatalf("Failed to read CSV record: %v", err)
			}
			chunk = append(chunk, record)
		}

		if len(chunk) == 0 {
			break
		}

		for _, record := range chunk {
			review := models.Review{
				ListingID:    string(record[0]),
				ID:           string(record[1]),
				Date:         string(record[2]),
				ReviewerID:   string(record[3]),
				ReviewerName: string(record[4]),
				Comments:     string(record[5]),
			}

			message, err := json.Marshal(review)
			if err != nil {
				log.Fatal(err)
			}

			hash := sha256.Sum256([]byte(message))
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
