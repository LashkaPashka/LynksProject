package kafka

import (
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	Writer *kafka.Writer
}

var ctx = context.Background()

func New(brokers []string, topic string) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}
	
	c := Client{}

	c.Writer = &kafka.Writer{
		Addr: kafka.TCP(brokers[0]),
		Topic: topic,
		Balancer: &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	return &c, nil
}

func (c *Client) Producer(data map[string]string) {
	var msgAll = []kafka.Message{}
	
	for key, val := range data {
		msg := kafka.Message{
			Key: []byte(key),
			Value: []byte(val),
		}
		msgAll = append(msgAll, msg)
	}

	err := c.Writer.WriteMessages(ctx, msgAll...)
	
	if err != nil {
		log.Println(err)
	}
}