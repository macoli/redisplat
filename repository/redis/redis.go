package redis

import (
	"context"
	"fmt"

	"github.com/macoli/redisplat/settings"

	"github.com/go-redis/redis/v8"
)

var (
	rc  *redis.Client
	ctx = context.Background()
)

// init redis
func Init(cfg *settings.RedisConfig) (err error) {
	rc = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rc.Ping(ctx).Result()
	return

}

func Close() {
	_ = rc.Close()
}
