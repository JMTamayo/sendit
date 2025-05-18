package services

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"email.assistant/src/config"
	"email.assistant/src/models"
)

type RedisService struct {
	client     *redis.Client
	streamName string
}

func NewRedisService(ctx context.Context) (*RedisService, *models.Error) {
	addr := fmt.Sprintf("%s:%d", config.Conf.GetRedisHost(), config.Conf.GetRedisPort())

	opts := &redis.Options{
		Addr:     addr,
		Username: config.Conf.GetRedisUsername(),
		Password: config.Conf.GetRedisPassword(),
		DB:       config.Conf.GetRedisDB(),
		Protocol: 2,
	}

	client := redis.NewClient(opts)
	_, err := client.Ping(ctx).Result()
	if err != nil {
		config.Log.Error(fmt.Sprintf("Error connecting to streams server: %v", err))
		return nil, models.NewError(fmt.Sprintf("Error connecting to streams server: %v", err))
	}

	return &RedisService{
		client:     client,
		streamName: config.Conf.GetStreamNameEmailQueue(),
	}, nil
}

func (s *RedisService) Produce(ctx context.Context, message any) (*string, *models.Error) {
	id, err := s.client.XAdd(ctx, &redis.XAddArgs{
		Stream: s.streamName,
		Values: message,
	}).Result()
	if err != nil {
		return nil, models.NewError(fmt.Sprintf("Error producing message to streams server: %v", err))
	}

	return &id, nil
}
