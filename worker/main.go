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
	topic := "coffee_orders"
	msgCnt := 0
	// 1: create a new consumer and start it
	worker, err := ConnectSonsumer([]string{"localhost:9092"})
	if err != nil {
		panic(err)
	}
	consumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}
	// 2: handle OS Signals -used to stop the process
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	// 3: create a GoRoutine to run the consumer/worker
	doneCh := make(chan struct{})
	go func() {
		for {
			select {
			case err := <-consumer.Errors():
				fmt.Println("err :", err)
			case msg := <-consumer.Messages():
				msgCnt++
				fmt.Printf("Recived Order Count %d |Topic (%s)|Message(%s)\n ", msgCnt, string(msg.Topic), string(msg.Value))
				order := string(msg.Value)
				fmt.Printf("Brewing cofee for order :%s\n", order)
			case <-sigchan:
				fmt.Println("Interrupt is detected ")
				doneCh <- struct{}{}
			}
		}

	}()
	<-doneCh
	fmt.Println("Processed", msgCnt, "messages")
	// 4: close the Consumer on exit
	if err := worker.Close(); err != nil {
		panic(err)
	}
}
func ConnectSonsumer(brokers []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	return sarama.NewConsumer(brokers, config)
}
