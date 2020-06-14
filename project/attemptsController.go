 package main

import (
	"net/http"
//    "html/template"
	"encoding/json"
//    "fmt"
)

var CreateAttempt = func(w http.ResponseWriter, r *http.Request) {

	attempt := &Attempt{}
	err := json.NewDecoder(r.Body).Decode(attempt) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	resp := attempt.Create() //Создать аккаунт
	Respond(w, resp)
}

var GetAnswer = func(w http.ResponseWriter, r *http.Request) {

	attempt := &AttemptAnswer{}
	err := json.NewDecoder(r.Body).Decode(attempt) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}
	 
	resp := attempt.Validate()
	Respond(w, resp)
}
