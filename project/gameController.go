package main

import (
	"net/http"
	"encoding/json"
//	"strconv"
//	"github.com/gorilla/mux"
 //   "fmt"
)

var CreateGame = func(w http.ResponseWriter, r *http.Request) {
	game := &Game{}
	err := json.NewDecoder(r.Body).Decode(game) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	
    if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

    game.Score = 0
	game.Attempts = 0
	resp := game.Create() //Создать аккаунт
	Respond(w, resp)
}

var GetGamesFor = func(w http.ResponseWriter, r *http.Request) {

	req := &GamesRequest{}
	
    err := json.NewDecoder(r.Body).Decode(req) //декодирует тело запроса в struct и завершается неудачно в случае ошибки
	
    if err != nil {
		Respond(w, Message(false, "Invalid request"))
		return
	}

	data := GetGames(req.Email)
    resp := Message(true, "success")
	resp["data"] = data
	Respond(w, resp)
}
