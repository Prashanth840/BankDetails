package kafka

import (
	"bankdetails/models"
	"bankdetails/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

var producer *kafka.Writer
var consumer *kafka.Reader

func ConnectKafka() {

	kafkaBroker := "localhost:9093"

	producer = &kafka.Writer{
		Addr:     kafka.TCP(kafkaBroker),
		Topic:    "transaction_topic",
		Balancer: &kafka.LeastBytes{},
	}

	fmt.Println("Kafka producer connected!")

	consumer = kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaBroker},
		GroupID:  "transaction-group",
		Topic:    "transaction_topic",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	fmt.Println("Kafka consumer connected!")
}

func PublishTransaction(transaction models.Transaction) error {
	message, err := json.Marshal(transaction)
	if err != nil {
		return fmt.Errorf("failed to marshal transaction: %v", err)
	}

	kafkaMessage := kafka.Message{
		Value: message,
	}

	err = producer.WriteMessages(context.Background(), kafkaMessage)
	if err != nil {
		return fmt.Errorf("failed to publish message to Kafka: %v", err)
	}

	fmt.Println("Transaction published to Kafka:", string(message))
	return nil
}

func ConsumeTransactions() {
	for {
		msg, err := consumer.ReadMessage(context.Background())
		if err != nil {
			log.Println("Error reading message:", err)
			continue
		}

		ProcessTransaction(msg)
	}
}

func ProcessTransaction(msg kafka.Message) {
	var transaction models.Transaction
	err := json.Unmarshal(msg.Value, &transaction)
	if err != nil {
		log.Println("Failed to unmarshal transaction:", err)
		return
	}

	account, err := services.GetAccountByID(transaction.AccountID)
	if err != nil {
		log.Println("Account not found:", err)
	}

	if transaction.TransactionType == "deposit" {
		account.Balance += transaction.Amount
	} else if transaction.TransactionType == "withdrawal" {
		if account.Balance < transaction.Amount {
			log.Println("Insufficient funds for withdrawal")
		}
		account.Balance -= transaction.Amount
	} else {
		log.Println("invalid Transaction Type")
	}

	err = services.UpdateAccount(account)
	if err != nil {
		log.Println("Failed to update account:", err)
	}

	transaction.Timestamp = time.Now()
	err = services.CreateTransaction(transaction)
	if err != nil {
		log.Println("Failed to log transaction:", err)
	}

	fmt.Println("Transaction processed successfully:", transaction)
}
