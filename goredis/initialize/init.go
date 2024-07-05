package initialize

import "github.com/redis/go-redis/v9"

var Rdb *redis.Client

func InitRedis() {
	//opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	//if err != nil {
	//	panic(err)
	//}
	//
	//rdb := redis.NewClient(opt)

	Rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // 密码
		DB:       0,  // 数据库
		PoolSize: 20, // 连接池大小
	})

	// 连接方式
	//	普通连接方式
	//  TLS 连接方式
	//  Redis Sentine 模式
	//  Redis Cluster 模式

}
