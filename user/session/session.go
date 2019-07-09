package session

import (
// "github.com/pengxianghu/v1-be/user/utils"
)

func InsertSession(u_id, s_id string) error {

	redisClient.Set(u_id, s_id, 0)

	return nil
}
