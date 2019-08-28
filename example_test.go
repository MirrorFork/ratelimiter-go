package ratelimiter_test

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"github.com/teambition/ratelimiter-go"
)

func ExampleRatelimiterGo() {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	limiter := ratelimiter.New(ratelimiter.Options{
		Client:   &redisClient{client},
		Max:      10,
		Duration: time.Second, // limit to 1000 requests in 1 minute.
	})

	userID := "user-123456"
	res, err := limiter.Get(userID)
	if err != nil {
		panic(err)
	}
	// fmt.Println(res.Reset) Reset time: 2016-10-11 21:17:53.362 +0800 CST
	fmt.Println(res.Total)
	fmt.Println(res.Remaining)
	fmt.Println(res.Duration)
	// Output:
	// 10
	// 9
	// 1s
}
