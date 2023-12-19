package lock

import (
	"commentsvr/config"
	"commentsvr/log"
	"fmt"
	"github.com/go-redsync/redsync/v4"
	pool "github.com/go-redsync/redsync/v4/redis"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var (
	rs      *redsync.Redsync
	rsOnce  sync.Once
	clients []*redis.Client

	lockExpiry = time.Second * 10
	retryDelay = time.Millisecond * 100
	tries      = 3
	//设置锁的过期时间
	//设计重试次数
	//设置重试间隔
	option = []redsync.Option{
		redsync.WithExpiry(lockExpiry),
		redsync.WithRetryDelay(retryDelay),
		redsync.WithTries(tries),
	}
	lockPrefix = "tiktok:lock:"
)

func initRedLock() {
	cfg := config.GetGlobalConfig().RedSyncConfig
	log.Infof("redLock init")
	pools := make([]pool.Pool, 0)
	for i, v := range cfg {
		addr := fmt.Sprintf("%s:%d", v.Host, v.Port)
		log.Infof("redLock %d cfg======%+v ", i, v)
		client := redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: v.PassWord,
			PoolSize: v.PoolSize,
		})
		_, err := client.Ping(context.Background()).Result()
		if err != nil {
			panic("redLock ping fatal")
		}
		clients = append(clients, client)
		pool := goredis.NewPool(client)
		pools = append(pools, pool)
	}
	rs = redsync.New(pools...)
}

func GetRedSync(name string) *redsync.Mutex {
	rsOnce.Do(initRedLock)
	return rs.NewMutex(lockPrefix+name, option...)
}

func CloseRedSync() {
	for _, v := range clients {
		v.Close()
	}
}

func Unlock(mutex *redsync.Mutex) {
	mutex.Unlock()
}
