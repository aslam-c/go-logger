package main

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	data,error:=json.Marshal(payload)
	if error!=nil{
		w.WriteHeader(400)
		return
	}
	w.Header().Add("Content-Type","application/json")
	w.WriteHeader(200)
	w.Write(data)
}


func HandleReadiness(w http.ResponseWriter,r *http.Request){
	RespondWithJSON(w,200,struct{}{})
}


func handleUserReq(w http.ResponseWriter,r *http.Request){
	RespondWithJSON(w,200,map[string]string{"hello":"user"})
	
}