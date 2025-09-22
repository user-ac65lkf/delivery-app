package main

import (
	kafkaReader "github.com/bakhtiyor-y/pushservice/internal/services/reader"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		kafkaReader.KafkaReader()
	}()
	wg.Wait()
}
