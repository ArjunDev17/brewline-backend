# **Kafka Coffee Order Processing System ☕**

This project demonstrates how to build a **Kafka-based order processing system** using **Golang**. It includes a producer, a consumer, and additional components to handle coffee orders efficiently.

---

## **Project Overview**
The project simulates a coffee shop's order management using **Apache Kafka** for message queuing and **Go** for implementation.  

### Components:
1. **Producer**:
   - Accepts orders via a REST API.
   - Sends orders as messages to a Kafka topic.
2. **Consumer**:
   - Listens to the Kafka topic.
   - Processes incoming coffee orders.

---

## **Getting Started**

### Prerequisites:
- [Go](https://golang.org/doc/install) (1.18+ recommended)
- [Apache Kafka](https://kafka.apache.org/quickstart) setup on `localhost:9092`
- Git installed

---

### Installation:
1. Clone the repository:
   ```bash
   git clone https://github.com/ArjunDev17/brewline-backend
   cd kafka-coffee-order-system
   ```
2. Switch to your desired branch:
   ```bash
   git branch
   git checkout <branch-name>
   ```
   - Available branches:
     - `main`: Main system setup.
     - `producer`: Contains producer code.
     - `consumers`: Contains consumer code.
     - `addtional_theory`: Extra resources and notes.

3. Install dependencies:
   ```bash
   go mod tidy
   ```

---

## **Usage**

### **Producer: Start the Server**
Run the producer to accept coffee orders via a REST API:
```bash
go run producer.go
```
- Endpoint: `POST http://localhost:3000/order`
- Example request:
  ```json
  {
      "customerName": "Kabir Singh",
      "coffeeType": "Milkyy"
  }
  ```

### **Consumer: Process Orders**
Run the consumer to process coffee orders:
```bash
go run consumer.go
```
- The consumer listens to the topic `coffee_orders` and logs processed orders.

---

## **Branches**
- **`main`**: Core system.
- **`producer`**: REST API for sending orders to Kafka.
- **`consumers`**: Processes messages from Kafka.
- **`addtional_theory`**: Notes or extra material for deeper learning.

---

## **Feedback**
Feel free to submit issues or pull requests. Your contributions are welcome!  

---
