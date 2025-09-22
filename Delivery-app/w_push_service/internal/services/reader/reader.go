package reader

import (
	"encoding/json"
	"fmt"
	"github.com/bakhtiyor-y/pushservice/internal/models"
	"github.com/bakhtiyor-y/pushservice/internal/service/pushV1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func KafkaReader() {
	// Настройка подключения к кафке
	config := &kafka.ConfigMap{
		"bootstrap.servers": "kafka:9092",             // адресс
		"group.id":          "my-consumer-group-push", // некая группа консьюмеров
		"auto.offset.reset": "earliest",               // с какого момента читать сообщения
	}

	// Создаем консьюмера
	consumer, err := kafka.NewConsumer(config)
	if err != nil {
		log.Println("failed to create consumer")
		//return
	}

	topics := []string{"mytopic"}

	err = consumer.SubscribeTopics(topics, nil)
	if err != nil {
		log.Println("failed to subscribe--> ", err)
		//return
	}

	log.Println("Жду сообщений----------->pushservice")

	var messageD models.MessageDetails

	for {
		msg1, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Println("failed to ReadMessage msg1")
		}
		msg := msg1.Value
		err = json.Unmarshal(msg, &messageD)
		if err != nil {
			log.Println("failed to json.unmarshal")
		}
		if err == nil {
			log.Println("received message", string(msg1.Value))
			pushV1.SendTelegramMessage(messageD)
		} else {
			fmt.Println("Error to read message pushservice")
		}
	}
}
