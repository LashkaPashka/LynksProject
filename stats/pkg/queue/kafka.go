package kafka

import (
	"Stats/configs"
	"Stats/internal/repository"
	"fmt"

	"Stats/internal/service"
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
		var mp = make(map[string]string)
		var msg kafka.Message
		var err error

		for i := 0; i < 2; i++ {
			msg, err = c.Reader.FetchMessage(context.Background())
			if err != nil {
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
				continue
			}
			mp[string(msg.Key)] = string(msg.Value)
		}
		
		err = c.Reader.CommitMessages(context.Background(), msg)
		if err != nil {
			log.Println(err)
		}
		
		c.service.CreateStat(mp)
	}
}



