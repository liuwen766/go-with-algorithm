package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {

	//kafkaPro()
	kafkaCon()

}

var (
	topic = "test-topic"
)

func kafkaCon() {
	consumer, _ := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "10.144.91.41",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	consumer.SubscribeTopics([]string{topic}, nil)

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			// The client will automatically try to recover from all errors.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	consumer.Close()
}

func kafkaPro() {
	config := kafka.ConfigMap{
		"bootstrap.servers": "10.144.91.41",
		//"bootstrap.servers": "10.139.7.8:9092",
		//"security.protocol": "SASL_PLAINTEXT",
		//"sasl.mechanism":    "PLAIN",
		//"sasl.username":     "admin",
		//"sasl.password":     "admin",
		//"linger.ms":         5,
		//"batch.size":        128,
		//"retries":           3,
	}

	producer, _ := kafka.NewProducer(&config)
	defer producer.Close()

	// Delivery report handler for produced messages
	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	// Produce messages to topic (asynchronously)
	fmt.Println("producer start send data to kafka!")
	for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
		producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)

		// Wait for message deliveries before shutting down
		producer.Flush(15 * 1000)
	}
}
