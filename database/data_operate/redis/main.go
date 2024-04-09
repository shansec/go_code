package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	val, err := rdb.Get(ctx, "key").Result()
	switch {
	case err == redis.Nil:
		fmt.Println("key不存在")
	case err != nil:
		fmt.Println("错误", err)
	case val == "":
		fmt.Println("值为空字符串")
	}
	//get := rdb.Get(ctx, "key")
	//if get.Err() != nil {
	//	fmt.Println("redis 命令执行错误", get.Err())
	//}
	//fmt.Println(get.Val())

}
