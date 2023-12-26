package cache

import (
	"commentsvr/config"
	"commentsvr/log"
	"fmt"
	redis "github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"sync"
	"time"
)

// 完成连接到redis并且返回conn对象
var (
	redisConn *redis.Client
	redisSync sync.Once
)

func initRedis() {
	cfg := config.GetGlobalConfig().RedisConfig
	log.Infof("redis cfg=====%+v", cfg)
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.PassWord,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}
	res, err := redisConn.Set(context.Background(), "abc", 100, 1*time.Second).Result()
	log.Infof("res====%v.err=======%v", res, err)
	_, err = redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("failed to ping redis ")
	}

}

func CloseRedis() {
	if redisConn != nil {
		redisConn.Close()
	}
}

func GetRdb() *redis.Client {
	redisSync.Do(initRedis)
	return redisConn
}
