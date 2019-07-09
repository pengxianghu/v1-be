package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/pengxianghu/v1-be/user/dbops"
	"github.com/pengxianghu/v1-be/user/defs"
	"github.com/pengxianghu/v1-be/user/utils"
	"github.com/pengxianghu/v1-be/user/session"
)

func registerUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bytes, _ := ioutil.ReadAll(r.Body)
	user := &defs.User{}

	if err := json.Unmarshal(bytes, user); err != nil {
		log.Printf("user register handler json unmarshal error.")
		sendErrResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	user.Id = utils.GenerateUserId()
	user.Pwd = utils.HashPwd(user.Pwd)
	if err := dbops.AddUser(user); err != nil {
		log.Printf("user register db error.")
		sendErrResponse(w, defs.ErroeDBError)
		return
	}

	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: user,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("user register handler json marshal error")
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}
	sendNormalResponse(w, http.StatusCreated, resp)
}

func loginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	u_name := ps.ByName("user_name")
	log.Printf("login name: %s", u_name)

	bytes, _ := ioutil.ReadAll(r.Body)
	user := &defs.User{}

	if err := json.Unmarshal(bytes, user); err != nil {
		log.Printf("user login handler json unmarshal error.")
		sendErrResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	db_user, err := dbops.GetUserCredential(u_name)
	if err != nil {
		log.Printf("user login handler dbops get pwd error: %s\n", err)
		sendErrResponse(w, defs.ErroeDBError)
		return
	}

	if db_user.Pwd != utils.HashPwd(user.Pwd) {
		log.Println("user login handler pwd auth fail.")
		sendErrResponse(w, defs.ErrorNotAuthUser)
		return
	}

	s_id, _ := utils.NewUUID()
	err = session.InsertSession("v1-be.user."+db_user.Id, s_id)
	if err != nil {
		log.Printf("user login handler insert session error.")
	}
	// return session
	http.SetCookie(w, &http.Cookie{
		Name:   "X-Session-Id",
		Value:  s_id,
		MaxAge: 1800,
	})

	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: user,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("user login handler json marshal error")
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}

	sendNormalResponse(w, http.StatusOK, resp)

}
