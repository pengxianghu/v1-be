/*
 * @Description:
 * @Autor: pengxianghu
 * @Date: 2019-08-10 11:24:58
 * @LastEditTime: 2019-08-17 20:28:57
 */
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
	// w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")

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

	router.POST("/logout", logouthandler)

	router.GET("/user/cookie", checkCookie)

	router.GET("/schedule/:id", getScheduleByIdHandler)

	router.GET("/schedules/user/:id", getScheduleByUserHandler)

	router.POST("/schedule", addScheduleHandler)

	router.PUT("/schedule", updateScheduleHandler)

	router.DELETE("/schedule/:s_id", deleteScheduleById)

	return router
}

func main() {
	r := registerHandler()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":25001", mh)
}
