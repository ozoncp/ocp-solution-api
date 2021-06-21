package producer

import (
	"github.com/Shopify/sarama"
)

type Producer interface {
	SendMessage(topic, message string) error
}

type producer struct {
	p sarama.SyncProducer
}

func New(brokers []string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	p, err := sarama.NewSyncProducer(brokers, config)

	return &producer{p: p}, err
}

func prepareMessage(topic, message string) *sarama.ProducerMessage {
	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(message),
	}
	return msg
}

func (p *producer) SendMessage(topic, message string) error {
	msg := prepareMessage(topic, message)
	_, _, err := p.p.SendMessage(msg)
	return err
}
