package config

import "time"

type Config struct {
	KafkaBroker        string
	KafkaTopic         string
	CSVFilePath        string
	KafkaChunkSize     int
	KafkaDelaySeconds  time.Duration
	KafkaConsumerGroup string
}

func NewConfig() *Config {
	return &Config{
		KafkaBroker:        "localhost:9092",
		KafkaTopic:         "test-topic3",
		CSVFilePath:        "reviews.csv",
		KafkaConsumerGroup: "review-consumer-group",
		KafkaChunkSize:     100,
		KafkaDelaySeconds:  1 * time.Second,
	}
}

func main() {
	config := NewConfig()
	println(config.KafkaBroker)
}
