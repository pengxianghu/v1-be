package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/pengxianghu/v1-be/dbops"
	"github.com/pengxianghu/v1-be/defs"
	"github.com/pengxianghu/v1-be/session"
	"github.com/pengxianghu/v1-be/utils"
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
		log.Printf("user register db error: %v", err)
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
		log.Printf("user register handler json marshal error.")
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
	user.Id = SESSION_PREFIX + db_user.Id

	s_id, _ := utils.NewUUID()
	log.Printf("session_id: %s", s_id)
	err = session.InsertSession(SESSION_PREFIX+db_user.Id, s_id)
	if err != nil {
		log.Printf("user login handler insert session error.")
		return
	}
	// return session
	// expire := time.Now().AddDate(0, 0, 1)
	s_c := http.Cookie{
		Name:     HEADER_FIELD_SESSION,
		Value:    s_id,
		Path:     "/",
		HttpOnly: true,
		// Expires: expire,
		MaxAge: 1800,
	}

	n_c := http.Cookie{
		Name:     HEADER_FIELD_UNAME,
		Value:    user.Id,
		Path:     "/",
		HttpOnly: false,
		// Expires: expire,
		MaxAge: 1800,
	}
	http.SetCookie(w, &s_c)
	http.SetCookie(w, &n_c)

	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: user,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("user login handler json marshal error: %v", err)
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}

	sendNormalResponse(w, http.StatusOK, resp)

}

func logouthandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	log.Printf("logout handler.")

	// 删除cookie
	n_c := http.Cookie{
		Name:     HEADER_FIELD_UNAME,
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1}
	http.SetCookie(w, &n_c)

	s_c := http.Cookie{
		Name:     HEADER_FIELD_SESSION,
		Path:     "/",
		HttpOnly: false,
		MaxAge:   -1}
	http.SetCookie(w, &s_c)

	w.Write([]byte("cookie已被删除."))
}

func addScheduleHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	bytes, _ := ioutil.ReadAll(r.Body)
	schedule := &defs.Schedule{}

	if err := json.Unmarshal(bytes, schedule); err != nil {
		log.Printf("add schedule handler json unmarshal error: %v", err)
		sendErrResponse(w, defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddSchedule(schedule); err != nil {
		log.Printf("add schedule db error: %v", err)
		sendErrResponse(w, defs.ErroeDBError)
		return
	}

	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: schedule,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("add schedule handler json marshal error: %v", err)
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}
	log.Println("add schedule success.")
	sendNormalResponse(w, http.StatusCreated, resp)

}

func getScheduleByUserHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	u_id := ps.ByName("id")

	s_list, err := dbops.GetScheduleByUser(u_id)

	if err != nil {
		log.Printf("get Schedule By user handler dbops error: %s\n", err)
		sendErrResponse(w, defs.ErroeDBError)
	}
	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: s_list,
	}
	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("get schedule by user handler json marshal error")
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}

	sendNormalResponse(w, http.StatusOK, resp)
}

func deleteScheduleById(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	s_id := ps.ByName("s_id")
	id, _ := strconv.Atoi(s_id)
	err := dbops.DeleteScheduleById(id)
	if err != nil {
		log.Printf("delete schedule dbops err: %v", err)
		sendErrResponse(w, defs.ErroeDBError)
		return
	}

	res := &defs.Result{
		Code: 0,
		Msg:  "success",
		Data: nil,
	}

	resp, err := json.Marshal(res)
	if err != nil {
		log.Printf("delete schedule by s_id handler json marshal error")
		sendErrResponse(w, defs.ErrorInternalFaults)
		return
	}

	sendNormalResponse(w, http.StatusOK, resp)
}

func checkCookie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	n_c, err := r.Cookie(HEADER_FIELD_UNAME)
	if err != nil {
		// w.Write([]byte("读取user name cookie失败: " + err.Error()))
		log.Printf("读取user name cookie失败")
	} else {
		data, _ := json.MarshalIndent(n_c, "", "\t")
		// w.Write([]byte("读取的user name cookie值: \n" + string(data)))
		log.Printf("读取user name cookie成功: %+v", string(data))
	}

	s_c, err := r.Cookie(HEADER_FIELD_SESSION)
	if err != nil {
		// w.Write([]byte("读取session cookie失败: " + err.Error()))
		log.Printf("读取session cookie失败")

	} else {
		data, _ := json.MarshalIndent(s_c, "", "\t")
		// w.Write([]byte("读取的session cookie值: \n" + string(data)))
		log.Printf("读取session cookie成功: %+v", string(data))
	}

}
