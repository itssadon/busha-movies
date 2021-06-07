package cache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/itssadon/busha-movies/collections"
)

var ctx = context.Background()

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) MoviesCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host,
		DB:       cache.db,
		Password: "vPAfLOsfockxKnJUBWODZEEl1NGhPuN7",
	})
}

func (cache *redisCache) Set(key string, value *collections.FilmCollection) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		// panic(err)
		log.Fatalln("Error marshaling movie collection for cache")
	}

	client.Set(ctx, key, string(json), cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *collections.FilmCollection {
	client := cache.getClient()

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		return nil
	}

	movie := collections.FilmCollection{}
	err = json.Unmarshal([]byte(val), &movie)
	if err != nil {
		panic(err)
	}

	return &movie
}
