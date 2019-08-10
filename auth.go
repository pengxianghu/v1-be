package main

import (
	"log"
	"net/http"
	"strings"
)

func validateUserSession(r *http.Request) bool {
	log.Printf("-- remote addr: %v, request url: %v --.", r.RemoteAddr, r.RequestURI)
	// return true
	if strings.Contains(r.URL.Path, "user") {
		log.Println("did not need auth")
		return true
	}
	s_c, err := r.Cookie(HEADER_FIELD_SESSION)
	if err != nil {
		log.Printf("get session cookie err: %v.", err)
		return false
	} else {
		log.Printf("session: %+v\n", s_c.Value)
	}

	n_c, err := r.Cookie(HEADER_FIELD_UNAME)
	if err != nil {
		log.Printf("get user name cookie err: %v.", err)
		return false
	} else {
		log.Printf("name: %+v\n", n_c.Value)
	}
	// data, _ := json.MarshalIndent(c, "", "\t")
	// log.Println("读取的cookie值: \n" + string(data))

	log.Println("auth passed")
	return true
}
