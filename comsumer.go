package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func Consumer() {
	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    TOPIC,
		MaxBytes: 1e3,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("Some error occur", err)
			continue
		}

		fmt.Println("Message:", string(m.Value))
	}
}

// func documentationVersion() {
// 	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", TOPIC, PARTITIONS)
// 	if err != nil {
// 		log.Fatal("Failed to Dial Leader:", err)
// 	}

// 	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
// 	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min 1MB max

// 	b := make([]byte, 10e3)
// 	fmt.Println("Start to read!")
// 	for {
// 		_, err := batch.Read(b)
// 		if err != nil {
// 			break
// 		}
// 		fmt.Println(string(b))
// 	}

// 	if err := batch.Close(); err != nil {
// 		log.Fatal("Failed to close batch:", err)
// 	}

// 	if err := conn.Close(); err != nil {
// 		log.Fatal("Failed to close connection:", err)
// 	}
// }
