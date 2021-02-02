package commands

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/Sho372/grc/client"
	"github.com/Sho372/grc/config"
	"github.com/go-redis/redis/v8"
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

func (a *App) Zadd(key string, score string, value string, period int, repeat int) {
	cxt := context.Background()

	if period == 0 {
		if repeat >= 1 {
			for i := 0; i < repeat; i++ {
				a.execZadd(cxt, key, strconv.Itoa(i))
			}
		}

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
					log.Println("[zadd]", t)
					if repeat >= 1 {
						for i := 0; i < repeat; i++ {
							a.execZadd(cxt, key, strconv.Itoa(i))
							fmt.Println(i)
						}
					}
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

func (a *App) execZadd(cxt context.Context, key string, index string) {

	unixTime := strconv.FormatInt(time.Now().Unix(), 10)
	dummyScore := unixTime + index
	dummyValue := "DUMMY VALUE: " + dummyScore

	_, err := a.Rdb.Rdb.Do(cxt, "zadd", key, dummyScore, dummyValue).Result()
	if err != nil {
		if err == redis.Nil {
			fmt.Printf("%s does not exists", key)
			return
		}
		log.Fatal(err)
	}
}
