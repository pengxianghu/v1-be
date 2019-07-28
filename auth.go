package main

import (
	"log"
	"net/http"
	"strings"
)

func validateUserSession(r *http.Request) bool {
	log.Println("----------")
	if strings.Contains(r.URL.Path, "user") {
		log.Println("did not need auth")
		return true
	}
	_, err := r.Cookie("X-Session-Id")
	if err != nil {
		log.Printf("get cookie err: %v, auth failed", err)
		return false
	}
	// data, _ := json.MarshalIndent(c, "", "\t")
	// log.Println("读取的cookie值: \n" + string(data))

	log.Println("auth passed")
	return true
}
