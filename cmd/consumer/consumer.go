package main

import (
	"flag"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	addr    = flag.String("addr", "nsqlookupd:4161", "NSQ lookupd addr")
	topic   = flag.String("topic", "", "NSQ Topic")
	channel = flag.String("channel", "", "NSQ Channel")
)

type MessageHandler struct{}

func (h *MessageHandler) HandleMessage(message *nsq.Message) error {
	log.Printf("Got a message: %s", string(message.Body))
	return nil
}

func main() {
	// pass the cli options
	if *topic == "" || *channel == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// configure a new consumer
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(*topic, *channel, config)
	if err != nil {
		log.Fatal(err)
	}

	// register our message handler with the consumer
	consumer.AddHandler(&MessageHandler{})

	// connect to NSQ and start receiving messages
	err = consumer.ConnectToNSQLookupd(*addr)
	if err != nil {
		log.Fatal(err)
	}

	// wait for the signal to exit
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	<-signalChan

	// disconnect
	consumer.Stop()
}
