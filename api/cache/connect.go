package cache

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()

func CreateRedisInstance(dbnumber int) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_URI"),
		Password: "",
		DB:       dbnumber,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("[REDIS] Error {%d} : %s", dbnumber, err.Error())
	}
	return rdb
}
