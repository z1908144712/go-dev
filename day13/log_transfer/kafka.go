package main

import (
	"github.com/Shopify/sarama"
)

var (
	kafkaClient *KafkaClient
)

type KafkaClient struct {
	client sarama.Consumer
}

func initKafka() (err error) {
	kafkaClient := KafkaClient{}
	consumer, err := sarama.NewConsumer([]string{appConfig.KafkaAddr}, nil)
	if err != nil {
		return
	}
	_, err = consumer.Partitions(appConfig.Topic)
	if err != nil {
		return
	}
	kafkaClient.client = consumer
	return
	// for _, partition := range partitionList {
	// 	pc, err := consumer.ConsumePartition(appConfig.Topic, int32(partition), sarama.OffsetOldest)
	// 	if err != nil {
	// 		return
	// 	}
	// 	defer pc.AsyncClose()
	// 	go func(pc sarama.PartitionConsumer) {
	// 		for msg := range pc.Messages() {
	// 			fmt.Printf("Partition:%d, Offset:%d, Key:%s, Value:%s\n", msg.Partition, msg.Offset, msg.Key, msg.Value)
	// 		}
	// 	}(pc)
	// }
}
