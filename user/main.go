package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func registerHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", registerUserHandler)

	router.POST("/user/:user_name", loginHandler)

	return router
}

// func main() {
// 	r := registerHandler()
// 	http.ListenAndServe(":25001", r)
// }

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	// validateUserSession(r)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("content-type", "application/json")

	m.r.ServeHTTP(w, r)
}

// func registerHandlers() *httprouter.Router {

// 	router := httprouter.New()

// 	router.POST("/user", CreateUser)

// 	router.POST("/login/:user_name", Login)

// 	return router

// }

func main() {
	r := registerHandler()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":25001", mh)
}
