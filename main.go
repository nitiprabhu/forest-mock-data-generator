package main

import (
	"awesomeProject/models"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
)

func main() {

	signals := make(chan os.Signal, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGKILL)

	ctx, cancel := context.WithCancel(context.Background())

	// go routine for getting signals asynchronously
	go func() {
		sig := <-signals
		fmt.Println("Got signal: ", sig)
		cancel()
	}()

	kafkaHost := os.Getenv("kafka.bootstrap");
	if kafkaHost == "" {
		kafkaHost = "localhost:9092"
	}

	kafkaTopic := os.Getenv("kafka.topic");
	if kafkaTopic == "" {
		kafkaTopic = "forestTopic"
	}

	bootstrapServers := strings.Split(kafkaHost, ",")
	delayMs, _ := strconv.Atoi(strconv.Itoa(1000))

	config := kafka.WriterConfig{
		Brokers:      bootstrapServers,
		Topic:        kafkaTopic,
		BatchTimeout: 1 * time.Millisecond}

	w := kafka.NewWriter(config)

	fmt.Println("Producer configuration: ", config)

	i := 1

	defer func() {
		err := w.Close()
		if err != nil {
			fmt.Println("Error closing producer: ", err)
			return
		}
		fmt.Println("Producer closed")
	}()

	for {
		message := models.GenerateRandomData()
		err := w.WriteMessages(ctx, kafka.Message{Value: []byte(message)})
		if err == nil {
			fmt.Println("Sent message: ", message)
		} else if err == context.Canceled {
			fmt.Println("Context canceled: ", err)
			break
		} else {
			fmt.Println("Error sending message: ", err)
		}
		i++

		time.Sleep(time.Duration(delayMs) * time.Millisecond)
	}
}