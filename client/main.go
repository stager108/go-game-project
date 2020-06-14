package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
  //  "net/url"
 //   "io/ioutil"
)

type ClientCredentials struct{
    email string
    token string
} 

type CurrentGame struct{
    game_number int
    current_attempt int
    Current_word string
} 

var credentials = &ClientCredentials{}

var current_game = &CurrentGame{}

func main() {

    router := mux.NewRouter()
    
    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })
        
    router.HandleFunc("/new", CreateAccount)

    router.HandleFunc("/login", Authenticate)
    
    router.HandleFunc("/new_game", NewGame)
    
    router.HandleFunc("/play", GetNewAttempt)
    router.HandleFunc("/send", SendAnswer)
    
    router.HandleFunc("/my_games", indexHandler)
	router.HandleFunc("/my_games/my_game/{id:[0-9]+}", gameHandler)
    
	port := os.Getenv("PORT") 
    
	if port == "" {
		port = "8001" 
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router)

	if err != nil {
		fmt.Print(err)
	}

}


 
 
