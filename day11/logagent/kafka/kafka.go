package kafka

import (
	"fmt"
	"go_dev/day11/logagent/commons"

	"github.com/astaxie/beego/logs"

	"github.com/Shopify/sarama"
)

var (
	client sarama.SyncProducer
)

func InitKafka(addr string) (err error) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	client, err = sarama.NewSyncProducer([]string{addr}, config)
	if err != nil {
		return
	}
	return
}

func SendToKafka(data *commons.TextMsg) {
	msg := &sarama.ProducerMessage{}
	msg.Topic = data.Topic
	msg.Value = sarama.StringEncoder(data.Text)
	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		logs.Error(err)
		return
	}
	fmt.Println("pid:", pid, "offset:", offset, "topic:", data.Topic, "text:", data.Text)
}
