package writer

import (
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func KafkaWriter() {
	// Настройка подключения к кафке
	config := &kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092", // адресс
	}

	// Создаем продюсера
	producer, err := kafka.NewProducer(config)
	if err != nil {
		log.Println("failed to create producer")
		return
	}

	producer.Flush(15 * 1000)
	topic := "mytopic"

	i := 0
	for {
		message := &kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &topic,
				Partition: kafka.PartitionAny,
			},
			Value: []byte(fmt.Sprintf("Message number: %d", i)),
		}

		err := producer.Produce(message, nil)
		if err != nil {
			log.Println("failed to send message")
			return
		}

		log.Println("Отправлено сообщение номер: ", i)
		i++

		time.Sleep(20 * time.Second)
	}
}
