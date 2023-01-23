// 封装 redis 客户端
// 必须先调用 Init 进行初始化
package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

// Nil 当键不存在时，有的命令会返回 Nil 错误
var Nil = redis.Nil

var client *redis.Client

// Init 初始化客户端
// 初始化时不会连接 redis
// addr 传空的话，默认连接 localhost:6379
func Init(addr string) {
	opt := &redis.Options{
		Addr: addr,
	}
	client = redis.NewClient(opt)
}

// Incr 将指定键的值加1，返回增加后的值
// 键不存在的话，先创建键并设置为0，后执行操作
// https://redis.io/commands/incr/
func Incr(ctx context.Context, key string) (int64, error) {
	return client.Incr(ctx, key).Result()
}

// Decr 将指定键的值减1，返回减小后的值
// 键不存在的话，先创建键并设置为0，后执行操作
// https://redis.io/commands/decr/
func Decr(ctx context.Context, key string) (int64, error) {
	return client.Decr(ctx, key).Result()
}

// Get 获取字符串键值
// 键不存在的话，返回 Nil 错误
// https://redis.io/commands/get/
func Get(ctx context.Context, key string) (string, error) {
	return client.Get(ctx, key).Result()
}
