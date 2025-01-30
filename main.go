package main

import (
	"github.com/bigdata1/kafka/consumer"
	"github.com/bigdata1/kafka/producer"
)

func main() {
	producer.Producer()
	consumer.Consumer()
}
