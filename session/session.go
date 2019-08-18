package session

import (
	"log"
	"time"
)

func InsertSession(u_id, s_id string) error {

	redisClient.Set(u_id, s_id, time.Second*1800)
	return nil
}

func GetSessionValue(key string) string {
	str, err := redisClient.Get(key).Result()
	if err != nil {
		log.Printf("get session error. key: %+v.", key)
		return ""
	}
	return str
}
