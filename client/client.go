package client

import (
	"github.com/Sho372/grc/config"
	"github.com/go-redis/redis/v8"
)

type Client struct {
	Rdb *redis.Client
}

func New(config *config.Config) *Client {
	r := redis.NewClient(&redis.Options{
		Addr:     config.Host +  ":" + config.Port ,
		Password: config.Password,
		DB:       config.Db,
	})
	return &Client{Rdb: r}
}
