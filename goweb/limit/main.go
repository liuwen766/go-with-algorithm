package main

import (
	"fmt"
	"go.uber.org/ratelimit"
	"time"
)

type common struct {
	a interface{}
	b string
	c int
}

func main() {

	fmt.Println("限流又称为流量控制（流控），通常是指限制到达系统的并发请求数。")
	fmt.Println("常用的限流策略：漏桶。它可以定义请求的速率，一个请求就是一滴水。")
	fmt.Println("假设速率是每秒滴下10滴水（即每秒处理10个请求）：")

	rl := ratelimit.New(10) // per second  每秒滴下10滴水
	prev := time.Now()
	for i := 0; i < 10; i++ {

		now := rl.Take()

		fmt.Println(i, now.Sub(prev))
		prev = now
	}

	fmt.Println("令牌桶其实和漏桶的原理类似，令牌桶按固定的速率往桶里放入令牌，并且只要能从桶里取出令牌就能通过，令牌桶支持突发流量的快速处理。")
	fmt.Println("在gin框架构建的项目中，我们可以将限流组件定义成中间件。如：func RateLimitMiddleware()")

	com := common{}
	s := fmt.Sprintf("%v", com.a)
	b := fmt.Sprintf("%v", com.a)
	if s != "" {
		fmt.Println(s, b)
	}
	fmt.Println(com.a, com.b, com.c)

}

//func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
//	bucket := ratelimit.NewBucket(fillInterval, cap)
//	return func(c *gin.Context) {
//		// 如果取不到令牌就中断本次请求返回 rate limit...
//		if bucket.TakeAvailable(1) < 1 {
//			c.String(http.StatusOK, "rate limit...")
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}
