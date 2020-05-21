package main

import (
	"fmt"

	"github.com/nsqio/go-nsq"
)

type ConsumerT struct{}

func main() {
	InitConsumer("test", "ch1", "127.0.0.1:4150")
	select {}
}

func (*ConsumerT) HandleMessage(msg *nsq.Message) error {
	fmt.Println("receive", msg.NSQDAddress, "message:", string(msg.Body))
	return nil
}

func InitConsumer(topic string, channel string, address string) {
	cfg := nsq.NewConfig()
	c, err := nsq.NewConsumer(topic, channel, cfg)
	if err != nil {
		panic(err)
	}
	c.AddHandler(&ConsumerT{})

	if err := c.ConnectToNSQD(address); err != nil {
		panic(err)
	}
}
