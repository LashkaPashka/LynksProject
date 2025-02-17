package kafka

import (
	"Lynks/stats/configs"
	"Lynks/stats/internal/repository"
	"Lynks/stats/internal/service"
	"context"
	"errors"
	"log"

	"github.com/segmentio/kafka-go"
)

type Client struct {
	Reader *kafka.Reader
	service *service.StatsService
	repo *repository.StatsRepository
}

var ctx = context.Background()

func New(brokers []string, topic, groupID string, conf *configs.Config) (*Client, error) {
	if len(brokers) == 0 || brokers[0] == "" || topic == "" || groupID == "" {
		return nil, errors.New("не указаны параметры подключения к Kafka")
	}
	
	c := Client{}

	c.Reader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokers,
		Topic: topic,
		GroupID: groupID,
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	c.repo, _ = repository.NewStatRepostiory(conf)
	c.service = service.NewStatService(c.repo)

	return &c, nil
}


func (c *Client) Consumer() {
	
	for {
		msg, err := c.Reader.FetchMessage(ctx)
		if err != nil {
			log.Println(err)
		}
		
		c.service.CreateStat(string(msg.Value))

		err = c.Reader.CommitMessages(ctx, msg)
		if err != nil {
			log.Println(err)
		}
	}
}
