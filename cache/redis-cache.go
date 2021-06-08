package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/itssadon/busha-movies/collections"
)

var ctx = context.Background()

type redisCache struct {
	host    string
	db      int
	pword   string
	expires time.Duration
}

func NewRedisCache(host string, db int, pword string, exp time.Duration) MoviesCache {
	return &redisCache{
		host:    host,
		db:      db,
		pword:   pword,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: cache.pword,
	})
}

func (cache *redisCache) Set(key string, value *collections.MovieCollection) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set(ctx, key, string(json), cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *collections.MovieCollection {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	movie := collections.MovieCollection{}
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}

	return &movie
}
