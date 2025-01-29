package config

import "time"

type Config struct {
	KafkaBroker       string
	KafkaTopic        string
	CSVFilePath       string
	KafkaChunkSize    int
	KafkaDelaySeconds time.Duration
}

func NewConfig() *Config {
	return &Config{
		KafkaBroker:       "localhost:9092",
		KafkaTopic:        "test-topic2",
		CSVFilePath:       "reviews.csv",
		KafkaChunkSize:    100,
		KafkaDelaySeconds: 1 * time.Second,
	}
}

func main() {
	config := NewConfig()
	println(config.KafkaBroker)
}
