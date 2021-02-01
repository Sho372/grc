package commands

import (
	"context"
	"fmt"
	"github.com/Sho372/grc/client"
	"github.com/Sho372/grc/config"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
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

func (a *App) Zadd(key string, score string, value string, period int) {
	cxt := context.Background()

	if period == 0 {
		val, err := a.Rdb.Rdb.Do(cxt, "zadd", key, score, value).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Printf("%s does not exists", key)
				return
			}
			log.Fatal(err)
		}
		log.Println(val)
	} else {
		ticker := time.NewTicker(time.Duration(a.Config.Interval) * time.Second)
		done := make(chan bool)

		log.Println("[zdd] started", ticker.C)

		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					_, err := a.Rdb.Rdb.Do(cxt, "zadd", key, score, value).Result()
					if err != nil {
						if err == redis.Nil {
							fmt.Printf("%s does not exists", key)
							return
						}
						log.Fatal(err)
					}
					log.Println("[zadd]", t)
				}
			}
		}()

		time.Sleep(time.Duration(period) * time.Second)
		log.Println("[zdd] stopped", ticker.C)
		ticker.Stop()
		done <- true
	}
}

func (a *App) Zrem(key string) {
}
