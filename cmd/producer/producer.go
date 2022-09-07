package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
)

var (
	addr    = flag.String("addr", "localhost:4150", "NSQ addr")
	topic   = flag.String("topic", "", "NSQ Topic")
	message = flag.String("message", "", "Message Body")
)

func main() {
	// parse the cli options
	flag.Parse()
	if *topic == "" || *message == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// configure a new Producer
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(*addr, config)
	if err != nil {
		log.Fatal(err)
	}

	// publish a message to producer
	err = producer.Publish(*topic, []byte(*message))
	if err != nil {
		log.Fatal(err)
	}

	producer.Stop()
}
