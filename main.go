package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// check session
	valid := validateUserSession(r)
	if !valid {
		// log.Printf("valid: %v", valid)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	m.r.ServeHTTP(w, r)
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
	m := middleWareHandler{}
	m.r = r
	return m
}

func registerHandler() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", registerUserHandler)

	router.POST("/user/:user_name", loginHandler)

	router.POST("/schedule", addScheduleHandler)

	router.GET("/schedule/:id", getScheduleByUserHandler)

	return router
}

func main() {
	r := registerHandler()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":25001", mh)
}