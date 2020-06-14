package main

import (
	"net/http"
	"strings"
	jwt "github.com/dgrijalva/jwt-go"
	"os"
	"context"
	"fmt"
)

var JwtAuthentication = func(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		notAuth := []string{"/",
                            "/api/user/new",
                            "/api/user/login"}
		requestPath := r.URL.Path 

		for _, value := range notAuth {
			if value == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}
		
		header := r.Header.Get("Authorization")
        fmt.Println(header)

		if header == "" {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, Message(false, "Authorization token is missed"))
			return
		}

		splitted := strings.Split(header, " ") 
        
		if len(splitted) != 2 {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, Message(false, "Strange authorization token"))
			return
		}

		token_string := splitted[1]
		token_struct := &Token{}

		token, err := jwt.ParseWithClaims(token_string, token_struct, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("token_password")), nil
		})

		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, Message(false, "Malformed authentication token"))
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusForbidden)
			w.Header().Add("Content-Type", "application/json")
			Respond(w, Message(false, "Invalid authorization token"))
			return
		}
		
		fmt.Sprintf("User %", token_struct.UserId)
		ctx := context.WithValue(r.Context(), "user", token_struct.UserId)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	});
}
