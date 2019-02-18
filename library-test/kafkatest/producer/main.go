package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}

	// Trap SIGINT to trigger a graceful shutdown.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	var (
		wg                          sync.WaitGroup
		enqueued, successes, errors int
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range producer.Successes() {
			successes++
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			log.Println(err)
			errors++
		}
	}()

	go func() {
	ProducerLoop:
		for {
			message := &sarama.ProducerMessage{Topic: "test", Value: sarama.ByteEncoder([]byte("testing 1"))}
			select {
			case producer.Input() <- message:
				enqueued++
			case <-signals:
				producer.AsyncClose() // Trigger a shutdown of the producer.
				break ProducerLoop
			}
		}
	}()

	go func() {
	ProducerLoop2:
		for {
			message := &sarama.ProducerMessage{Topic: "test", Value: sarama.ByteEncoder([]byte("testing 2"))}
			select {
			case producer.Input() <- message:
				enqueued++
			case <-signals:
				producer.AsyncClose() // Trigger a shutdown of the producer.
				break ProducerLoop2
			}
		}
	}()

	time.Sleep(1 * time.Second)
	// wg.Wait()

	log.Printf("Successfully enqueued: %d, produced: %d; errors: %d\n", enqueued, successes, errors)
}
