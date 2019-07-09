package main

import (
	"encoding/json"
	"net/http"

	"github.com/pengxianghu/v1-be/user/defs"
)

func sendErrResponse(w http.ResponseWriter, err defs.Err) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.ErrorCode)
	ubody, _ := json.Marshal(err.Error)
	w.Write(ubody)
}

func sendNormalResponse(w http.ResponseWriter, sc int, resp []byte) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(sc)
	w.Write(resp)
}
