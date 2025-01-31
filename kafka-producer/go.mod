module github.com/bigdata1/kafka-producer

go 1.23.5

require (
	github.com/bigdata1/config v0.0.0-00010101000000-000000000000
	github.com/bigdata1/models v0.0.0-00010101000000-000000000000
	github.com/confluentinc/confluent-kafka-go v1.9.2
)

replace github.com/bigdata1/config => ./config

replace github.com/bigdata1/models => ./models
