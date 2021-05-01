package main

import (
	"net"
	"strconv"

	"github.com/segmentio/kafka-go"
)

func CreateTopic() {
	topic := "first-topic"
	partition := 2

	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}

	var controllerConn *kafka.Conn
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := kafka.TopicConfig{
		Topic:             topic,
		NumPartitions:     partition,
		ReplicationFactor: 1,
	}

	err = controllerConn.CreateTopics(topicConfigs)
	if err != nil {
		panic(err.Error())
	}
}
