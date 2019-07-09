package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"strconv"
	"time"
)

const salt = "./a2;%78@z,al[0e]"

func GenerateUserId() string {
	f := time.Now().Unix()
	b := rand.Int63n(1000)
	return strconv.FormatInt(f, 10) + strconv.FormatInt(b, 10)
}

func HashPwd(pwd string) string {
	data := []byte(pwd + salt)
	b := md5.Sum(data)
	md5str := fmt.Sprintf("%x", b)
	return md5str
}

// func GenerateSessionId() {

// }

func NewUUID() (string, error) {
	u, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return u.String(), nil
}
