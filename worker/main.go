package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

/*
the consumer listen to the cofee orders topic and processes the incoming orders
       1: create a new consumer and start it
       2: handle OS Signals -used to stop the process .
       3: create a GoRoutine to run the consumer/worker
       4: close the Consumer in exit
*/

func main() {
	topic := "coffee_orders"                                   // The topic to listen to.
	msgCnt := 0                                                // Counter for processed messages.
	worker, err := ConnectSonsumer([]string{"localhost:9092"}) // Connect to Kafka.
	if err != nil {
		panic(err)
	} // Stop if connection fails.
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest) // Start consuming messages.
	if err != nil {
		panic(err)
	} // Stop if partition fails.
	sigchan := make(chan os.Signal, 1)                      // Create channel for OS signals.
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM) // Notify on interrupt signals.
	doneCh := make(chan struct{})                           // Channel to signal when to stop.
	go func() {                                             // Goroutine to process messages.
		for {
			select {
			case err := <-consumer.Errors(): // Handle consumer errors.
				fmt.Println("err :", err)
			case msg := <-consumer.Messages(): // Handle new messages.
				msgCnt++
				fmt.Printf("Recived Order Count %d |Topic (%s)|Message(%s)\n ", msgCnt, string(msg.Topic), string(msg.Value))
				order := string(msg.Value)
				fmt.Printf("Brewing cofee for order :%s\n", order)
			case <-sigchan: // Handle shutdown signal.
				fmt.Println("Interrupt is detected ")
				doneCh <- struct{}{}
			}
		}
	}()
	<-doneCh                                     // Wait until the goroutine signals completion.
	fmt.Println("Processed", msgCnt, "messages") // Log the total messages processed.
	if err := worker.Close(); err != nil {       // Close the Kafka consumer.
		panic(err)
	}
}

func ConnectSonsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}

// T9w#3k!2D@7mQ8
