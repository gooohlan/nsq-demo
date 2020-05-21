package main

import (
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
	"time"
)

func TestNSQ1(t *testing.T) {
	NSQDsAddrs := []string{"127.0.0.1:4150", "127.0.0.1:4152"}
	go consumer1(NSQDsAddrs)
	go produce1()
	go produce2()
	time.Sleep(30 * time.Second)
}

func produce1() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := producer.Publish("test", []byte("x")); err != nil {
		log.Fatal("publish error:" + err.Error())
	}
	if err := producer.Publish("test", []byte("y")); err != nil {
		log.Fatal("publist error:" + err.Error())
	}
}

func produce2() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4152"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := producer.Publish("test", []byte("z")); err != nil {
		log.Fatal("publist error:" + err.Error())
	}
}

func consumer1(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor01", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C1")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C1")
	}
	<-consumer.StopChan
}
func TestNSQ2(t *testing.T) {
	NSQDsAddrs := []string{"127.0.0.1:4150"}
	go consumer1(NSQDsAddrs)
	go consumer2(NSQDsAddrs)
	go consumer3(NSQDsAddrs)
	go produce3()
	time.Sleep(5 * time.Second)
}

func produce3() {
	cfg := nsq.NewConfig()
	nsqdAddr := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqdAddr, cfg)
	if err != nil {
		log.Fatal(err)
	}
	if err := producer.Publish("test", []byte("x")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
	if err := producer.Publish("test", []byte("y")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
	if err := producer.Publish("test", []byte("z")); err != nil {
		log.Fatal("publish error: " + err.Error())
	}
}

func consumer2(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor02", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C2")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C2")
	}
	<-consumer.StopChan
}

func consumer3(NSQDsAddrs []string) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "sensor02", cfg)
	if err != nil {
		log.Fatal(err)
	}
	consumer.AddHandler(nsq.HandlerFunc(
		func(message *nsq.Message) error {
			log.Println(string(message.Body) + " C3")
			return nil
		}))
	if err := consumer.ConnectToNSQDs(NSQDsAddrs); err != nil {
		log.Fatal(err, " C3")
	}
	<-consumer.StopChan
}
