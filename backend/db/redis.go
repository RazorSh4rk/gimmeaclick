package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client
var ctx = context.Background()
var storageTime time.Duration = 24 * 3 * time.Hour

func Connect() {
	raw := os.Getenv("REDIS_URL")

	uri, err := redis.ParseURL(raw)

	if err != nil {
		log.Fatal(err)
	}

	rdb = redis.NewClient(uri)
}

func Put(key string, value string) error {
	return rdb.Set(ctx, key, value, storageTime).Err()
}

func Get(key string) (string, error) {
	val, err := rdb.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return val, nil
}

func GetAll() (map[string]string, error) {
	collection := make(map[string]string)

	keys, err := rdb.Keys(ctx, "*").Result()

	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		rec, err := Get(key)
		if err != nil {
			rec = ""
		}

		collection[key] = rec
	}

	return collection, nil
}
