package main

import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
)


func main() {

	router := mux.NewRouter()
	router.Use(JwtAuthentication) // добавляем middleware проверки JWT-токена
    
    router.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome to my website!")
    })
        
    router.HandleFunc("/api/user/new", CreateAccount).Methods("POST")

    router.HandleFunc("/api/user/login", Authenticate).Methods("POST")

    router.HandleFunc("/api/me/all_games", GetGamesFor).Methods("POST")
    
    router.HandleFunc("/api/me/new_game", CreateGame).Methods("POST")
    
    router.HandleFunc("/api/me/new_attempt", CreateAttempt).Methods("POST")
    
    router.HandleFunc("/api/me/send_answer", GetAnswer).Methods("POST")

	port := os.Getenv("PORT") 
    //Получить порт из файла .env; мы не указали порт, поэтому при локальном тестировании должна возвращаться пустая строка
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Запустите приложение, посетите localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
