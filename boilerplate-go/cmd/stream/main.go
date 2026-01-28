package main

import (
	"boilerplate-go/internal/application/config"
	"boilerplate-go/internal/infra/kafka"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	pkgKafka "github.com/segmentio/kafka-go"
)

func main() {
	envs := config.LoadEnvs()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	brokers := strings.Split(envs.KafkaBrokers, ",")
	subscriber := kafka.NewKafkaSubscriber(brokers, envs.KafkaTopic, envs.KafkaGroupID)
	defer subscriber.Close()

	fmt.Println("Stream processor started")

	err := subscriber.Subscribe(ctx, func(ctx context.Context, msg pkgKafka.Message) error {
		// Business logic for processing messages goes here
		fmt.Printf("Processing message: %s\n", string(msg.Value))
		return nil
	})

	if err != nil {
		log.Fatalf("Stream processor failed: %v", err)
	}

	fmt.Println("Stream processor stopped gracefully")
}
