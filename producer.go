package main

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

func Producer() {
	conf := kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   TOPIC,
	}

	producer := kafka.NewWriter(conf)

	for {
		err := producer.WriteMessages(context.Background(), kafka.Message{Value: []byte("First Message")})
		if err != nil {
			log.Fatal(err.Error())
		}
		time.Sleep(time.Second * 1)
	}
}

// func documentationVersion() {
// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "first-topic", 1)
// 	if err != nil {
// 		log.Fatal("Failed to dial leader:", err)
// 	}

// 	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
// 	_, err = conn.WriteMessages(
// 		kafka.Message{Value: []byte("First Message")},
// 		kafka.Message{Value: []byte("Second Message")},
// 		kafka.Message{Value: []byte("Third Message")},
// 	)
// 	if err != nil {
// 		log.Fatal("Failed to write messages:", err)
// 	}

// 	if err := conn.Close(); err != nil {
// 		log.Fatal("Failed to close writer:", err)
// 	}
// }
