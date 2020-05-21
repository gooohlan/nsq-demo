package main

import (
	"bufio"
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"os"
	"strings"
)

func main() {
	cfg := nsq.NewConfig()
	nsqd := "127.0.0.1:4150"
	producer, err := nsq.NewProducer(nsqd, cfg)
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> topic: ")
		topic, _ := reader.ReadString('\n')
		topic = strings.Replace(topic, "\n", "", -1)
		fmt.Print("-> message: ")
		message, _ := reader.ReadString('\n')
		message = strings.Replace(message, "\n", "", -1)
		fmt.Println("消息发送中\n")

		if err := producer.Publish(topic, []byte(message)); err != nil {
			log.Fatal("publish error:" + err.Error())
		}

	}
}
