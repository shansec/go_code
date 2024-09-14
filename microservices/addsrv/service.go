package main

import (
	"context"
	"errors"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

const MaxLen = 10

var (
	// ErrTwoZeroes  Sum方法的业务规则不能对两个0求和
	ErrTwoZeroes = errors.New("can't sum two zeroes")

	// ErrIntOverflow Sum参数越界
	ErrIntOverflow = errors.New("integer overflow")

	// ErrTwoEmptyStrings Concat方法业务规则规定参数不能是两个空字符串.
	ErrTwoEmptyStrings = errors.New("can't concat two empty strings")

	// ErrMaxSizeExceeded Concat方法的参数超出范围
	ErrMaxSizeExceeded = errors.New("result exceeds maximum size")
)

type AddService interface {
	Sum(ctx context.Context, a, b int) (int, error)
	Concat(ctx context.Context, a, b string) (string, error)
}

//type addService struct{}
//
//func (s addService) Sum(_ context.Context, a, b int) (int, error) {
//	if a == 0 && b == 0 {
//		return 0, ErrTwoZeroes
//	}
//	if (b > 0 && a > (math.MaxInt-b)) || (b < 0 && a < (math.MinInt-b)) {
//		return 0, ErrIntOverflow
//	}
//	return a + b, nil
//}
//
//func (s addService) Concat(_ context.Context, a, b string) (string, error) {
//	if a == "" && b == "" {
//		return "", ErrTwoEmptyStrings
//	}
//	if len(a)+len(b) > MaxLen {
//		return "", ErrMaxSizeExceeded
//	}
//	return a + b, nil
//}
//
//// NewService 创建一个add service
//func NewService() AddService {
//	return &addService{}
//}

// 应用层使用中间件添加日志
type logMiddleware struct {
	logger log.Logger
	next   AddService
}

func (l logMiddleware) Sum(ctx context.Context, a, b int) (ret int, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "Sum",
			"a", a,
			"b", b,
			"ret", ret,
			"err", err,
			"time", time.Since(begin),
		)
	}(time.Now())
	ret, err = l.next.Sum(ctx, a, b)
	return
}

func (l logMiddleware) Concat(ctx context.Context, a, b string) (ret string, err error) {
	defer func(begin time.Time) {
		l.logger.Log(
			"method", "Concat",
			"a", a,
			"b", b,
			"ret", ret,
			"err", err,
			"time", time.Since(begin),
		)
	}(time.Now())
	ret, err = l.next.Concat(ctx, a, b)
	return
}

func NewLogMiddlewareService() AddService {
	return &logMiddleware{}
}

func NewLogMiddleware(logger log.Logger, next AddService) AddService {
	return &logMiddleware{
		logger: logger,
		next:   next,
	}
}

// 调用其它服务
type withTrimMiddleware struct {
	next        AddService
	trimService endpoint.Endpoint
}
