package sql_demo

import (
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis/v8"
)

func TestDoSomethingWithRedis(t *testing.T) {
	// mock 一个 redis server
	s, err := miniredis.Run()
	if err != nil {
		panic(err)
	}
	defer s.Close()

	// 准备数据
	s.Set("may", "五月")
	s.SAdd(KeyValidWebsite, "may")

	// 连接 mock 的 redis server
	rdb := redis.NewClient(&redis.Options{
		Addr: s.Addr(),
	})

	ok := DoSomethingWithRedis(rdb, "may")
	if !ok {
		t.Fatal()
	}

	// 可以手动检查redis中的值是否复合预期
	//if got, err := s.Get("blog"); err != nil || got != "https://liwenzhou.com" {
	//	t.Fatalf("'blog' has the wrong value")
	//}
	// 也可以使用帮助工具检查
	s.CheckGet(t, "blog", "https://liwenzhou.com")

	// 过期检查
	s.FastForward(5 * time.Second) // 快进5秒
	if s.Exists("blog") {
		t.Fatal("'blog' should not have existed anymore")
	}
}
