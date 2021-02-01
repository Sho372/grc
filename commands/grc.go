package commands

import (
	"context"
	"fmt"
	"github.com/Sho372/grc/client"
	"github.com/Sho372/grc/config"
	"github.com/go-redis/redis/v8"
	"log"
)

type App struct {
	Rdb    *client.Client
	Config *config.Config
}

func New() (*App, error) {
	// Load a config
	c, err := config.Load()
	if err != nil {
		return nil, err
	}
	// Create a redis client
	r := client.New(c)

	a := &App{
		Rdb:    r,
		Config: c,
	}
	return a, nil
}

func (a *App) Zadd(key string, score string, value string) {
	cxt := context.Background()
	val, err := a.Rdb.Rdb.Do(cxt, "zadd", key, score, value).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("%s does not exists", key)
			return
		}
		log.Fatal(err)
	}
	log.Println(val)
}

func (a *App) Zrem(key string) {
}
