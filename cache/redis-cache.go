package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gitlab.com/pragmaticreviews/golang-mux-api/entity"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) PostCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

type keployRedisClient struct {
	redis.Client
}

func (krc *keployRedisClient) Process(cmd redis.Cmder) error {
	fmt.Println("\ninto keploys's redis Process method")
	return nil
}

func (krc *keployRedisClient) ProcessContext(ctx context.Context, cmd redis.Cmder) error {
	fmt.Println("\ninto keploys's redis ProcessContext method")
	return nil
}

func (krc *keployRedisClient) Do(ctx context.Context, args ...interface{}) *redis.Cmd {
	fmt.Println("\ninto keploys's redis Do method")
	return krc.Client.Do(ctx, args...)
}

func (krc *keployRedisClient) Get(ctx context.Context, key string) *redis.StringCmd {
	fmt.Println("\ninto keploys's redis Get method")
	return krc.Client.Get(ctx, key)
}

func (krc *keployRedisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	fmt.Println("\ninto keploys's redis Get method")
	return krc.Client.Set(ctx, key, value, expiration)
}

func (cache *redisCache) getClient() *keployRedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     cache.host,
		Password: "",
		DB:       cache.db,
	})
	return &keployRedisClient{Client: *client}
}

func (cache *redisCache) Set(key string, post *entity.Post) {
	client := cache.getClient()

	// client.Process(nil)
	// serialize Post object to JSON
	json, err := json.Marshal(post)
	if err != nil {
		panic(err)
	}

	client.Set(context.Background(), key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *entity.Post {
	client := cache.getClient()

	// val, err := client.Do("get", key).Result()

	val, err := client.Get(context.Background(), key).Result()
	if err != nil {
		return nil
	}

	post := entity.Post{}
	err = json.Unmarshal([]byte(val), &post)
	if err != nil {
		panic(err)
	}
	return &post
}
