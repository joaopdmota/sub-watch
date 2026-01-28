package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

type KafkaSubscriber struct {
	reader *kafka.Reader
}

func NewKafkaSubscriber(brokers []string, topic, groupID string) *KafkaSubscriber {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	return &KafkaSubscriber{
		reader: reader,
	}
}

func (s *KafkaSubscriber) Subscribe(ctx context.Context, handler func(ctx context.Context, msg kafka.Message) error) error {
	fmt.Printf("Starting Kafka subscriber for topic: %s...\n", s.reader.Config().Topic)
	for {
		m, err := s.reader.FetchMessage(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return nil
			}
			return fmt.Errorf("failed to fetch message: %w", err)
		}

		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		if err := handler(ctx, m); err != nil {
			log.Printf("error handling message: %v", err)
			continue // Depending on requirements, we might want to retry or commit anyway
		}

		if err := s.reader.CommitMessages(ctx, m); err != nil {
			return fmt.Errorf("failed to commit message: %w", err)
		}
	}
}

func (s *KafkaSubscriber) Close() error {
	fmt.Println("Closing Kafka subscriber...")
	return s.reader.Close()
}
