/*
 * @Description:
 * @Autor: pengxianghu
 * @Date: 2019-08-10 11:24:58
 * @LastEditTime: 2019-08-17 21:57:35
 */
package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/pengxianghu/v1-be/session"
)

func validateUserSession(r *http.Request) bool {
	log.Printf("-- remote addr: %v, request url: %v --.", r.RemoteAddr, r.RequestURI)
	if strings.Index(r.URL.Path, "/user") == 0{
		log.Println("did not need auth")
		return true
	}
	// return true
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

	s_val := session.GetSessionValue(n_c.Value)

	if s_val != s_c.Value {
		log.Println("auth failed.")
		return false
	}

	log.Println("auth passed.")
	session.InsertSession(n_c.Value, s_c.Value)
	return true
}
