package session

import (
	"github.com/go-redis/redis"
)

var redisClient *redis.Client

// func init() {
// 	redisClient = redis.NewClient(&redis.Options{
// 		Addr:     "119.23.70.24:25004",
// 		Password: "",
// 		DB:       0,
// 	})

// 	_, err := redisClient.Ping().Result()
// 	if err != nil {
// 		panic(err)
// 	}
// }


func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})

	_, err := redisClient.Ping().Result()
	if err != nil {
		panic(err)
	}
}
