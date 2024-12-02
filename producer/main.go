package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "github.com/Shopify/sarama"
	"github.com/IBM/sarama"
)

type Order struct {
	CustomerName string `json:"customerName"`
	CofeeType    string `json:"cofeeType"`
}

func main() {
	http.HandleFunc("/order", placeOrder)
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// 3.2
func ConnectProducer(brokers []string) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	return sarama.NewSyncProducer(brokers, config)
}

// 3.1
func PushOrderToQueue(topic string, message []byte) error {
	brokers := []string{"localhost:9092"}
	//create Conncetion
	producer, err := ConnectProducer(brokers)
	if err != nil {
		return err
	}
	defer producer.Close()

	//Create a new message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	fmt.Printf("Order is stored in (%s)/patition(%d)/offset(%d)\n ", topic, partition, offset)
	return nil

}
func placeOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request Method", http.StatusMethodNotAllowed)
		return
	}
	//to impliment producer ,,that write events to kafka we need to folloow this steps :

	/*
		1: parse request body into order
		2: convert body into bytes
		3: send the bytes to kafka
		4: respond back to the user
	*/

	//1: parse request body into order
	order := new(Order)
	if err := json.NewDecoder(r.Body).Decode(order); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//2: convert body into bytes
	orderInBytes, err := json.Marshal(order)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 3: send the bytes to kafka
	err = PushOrderToQueue("cofee_orders", orderInBytes)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 4: respond back to the user
	response := map[string]interface{}{
		"Success": true,
		"msg":     "Order for " + order.CustomerName + " placed Successfully !",
	}
	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(response); err != nil {
		log.Println(err)
		http.Error(w, "error placing order", http.StatusInternalServerError)
		return
	}

}
