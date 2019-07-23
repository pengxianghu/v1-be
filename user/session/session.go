package session

import (
	"time"
)

func InsertSession(u_id, s_id string) error {

	redisClient.Set(u_id, s_id, time.Second * 1800)

	return nil
}
