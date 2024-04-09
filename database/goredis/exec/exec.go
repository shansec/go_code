package exec

import (
	"context"
	"fmt"
	"go_redis_study/initialize"
	"time"
)

func DoCommand() {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 执行命令获取结果
	val, err := initialize.Rdb.Get(ctx, "key").Result()
	fmt.Println(val, err)

	// 先获取到命令对象
	cmder := initialize.Rdb.Get(ctx, "key")
	fmt.Println(cmder.Val())
	fmt.Println(cmder.Err())

	// 直接执行命令获取错误
	err = initialize.Rdb.Set(ctx, "key", 10, time.Hour).Err()

	// 直接执行命令获取值
	value := initialize.Rdb.Get(ctx, "key").Val()
	fmt.Println(value)
}
