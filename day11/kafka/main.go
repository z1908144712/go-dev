package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func main() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()
	msg := &sarama.ProducerMessage{}
	msg.Topic = "nginx_log"
	msg.Value = sarama.StringEncoder("this is a test")
	for {
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("pid:", pid, "offset:", offset)
	}

}
