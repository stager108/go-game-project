package main

import (
	"net/http"
//    "html/template"
	"encoding/json"
//    "fmt"
)

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {

	account := &Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	resp := account.Create() //Создать аккаунт
	Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {

	account := &Account{}
	err := json.NewDecoder(r.Body).Decode(account) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	resp := Login(account.Email, account.Password)
	Respond(w, resp)
}
