package conn

import (
	"context"
	"log/slog"
	"time"

	"ddr13/riddlercore/internal/cache"
	"ddr13/riddlercore/internal/config"

	"github.com/go-redis/redis/v8"
)

var defaultCache cache.Cache
var redisClient *redis.Client

// DefaultRedis return connected redis default client
func DefaultRedis() *redis.Client {
	return redisClient
}

// ConnectRedis ...
func ConnectRedis(cfg *config.RedisConfig) error {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
	})
	defaultCache = cache.NewRedis(rdb)
	redisClient = rdb
	return rdb.Ping(rdb.Context()).Err()
}

// ConnectDefaultRedis connect with default configurations
func ConnectDefaultRedis() error {
	cfg := config.Redis()
	err := ConnectRedis(cfg)
	// run a background process to ping and establish connection
	go func() {
		for {
			if err := defaultCache.Ping(context.Background()); err != nil {
				slog.Warn("cache: ping error:", err)
				if err := ConnectRedis(cfg); err != nil {
					slog.Warn("cache:failed to reconnect:", err)
				}
			}
			time.Sleep(5 * time.Second)
		}
	}()
	return err
}

// DefaultCache return default connected cache
func DefaultCache() cache.Cache {
	return defaultCache
}
