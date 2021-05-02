package main

import (
	"log"
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		log.Fatal("Failed to Dial:", err)
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		log.Fatal("Failed to create controller:", err)
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		log.Fatal("Failed to Admin Dial:", err)
	}
	defer controllerConn.Close()

	topicConfigs := kafka.TopicConfig{
		Topic:             TOPIC,
		NumPartitions:     PARTITIONS,
		ReplicationFactor: 1,
	}

	err = controllerConn.CreateTopics(topicConfigs)
	if err != nil {
		log.Fatal("Failed to desconnect:", err)
	}
}
