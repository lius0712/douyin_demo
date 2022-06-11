package cache

import (
	"context"
	"github.com/RaymondCode/simple-demo/repository"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx = context.Background()

func Exists(key string) bool {
	return repository.RDB.Exists(ctx, key).Val() == 1
}

func Set(key string, value interface{}, expiration time.Duration) {
	if Exists(key) {
		repository.RDB.Expire(ctx, key, expiration)
		return
	}
	repository.RDB.Set(ctx, key, value, expiration)
}

func Get(key string) *redis.StringCmd {
	return repository.RDB.Get(ctx, key)
}

func Incr(key string) {
	repository.RDB.Incr(ctx, key)
}

//向集合添加一个或多个成员
func SAdd(key string, members interface{}) {
	repository.RDB.SAdd(ctx, key, members)
}

//移除集合中一个或多个成员
func SRem(key string, members interface{}) {
	repository.RDB.SRem(ctx, key, members)
}
