package main

import (
	"net/http"
    "html/template"
	"encoding/json"
  //  "strconv"
    "fmt"
    "io/ioutil"
    "bytes"
    "github.com/tidwall/gjson"
)


var Authenticate = func(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        request, err := json.Marshal(map[string]string{
                    "email":   r.FormValue("email"),
                    "password" : r.FormValue("password")})
        
        if err != nil {
            Respond(w, Message(false, "Json encode error!"))
            return
        }
        
        resp, err := http.Post("http://project:8000/api/user/login",
            "application/json", bytes.NewBuffer(request))
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        
        if err != nil {
            Respond(w, Message(false, "Json encode error!"))
            return
        }
        
        value := gjson.Get(string(body), "status").Bool()
        fmt.Println(value)
	    if value {
            credentials.email = gjson.Get(string(body), "account.email").String()
            credentials.token = gjson.Get(string(body), "account.token").String()
        }
        
        message := gjson.Get(string(body), "message").String()
        Respond(w, Message(false, message))
        
    }
}

var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("create.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        request, err := json.Marshal(map[string]string{
                    "email":   r.FormValue("email"),
                    "password" : r.FormValue("password")})
        
        if err != nil {
            Respond(w, Message(false, "Json encode error!"))
            return
        }
        
        resp, err := http.Post("http://project:8000/api/user/new",
            "application/json", bytes.NewBuffer(request))
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        
        if err != nil {
            Respond(w, Message(false, "Json encode error!"))
            return
        }
        
        message := gjson.Get(string(body), "message").String()
        Respond(w, Message(false, message))
        
    }
}

var NewGame = func(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("new_game.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        request, err := json.Marshal(map[string]string{
                    "email":   credentials.email})
        
        if err != nil {
            Respond(w, Message(false, "Json decode error!"))
            return
        }
        
        client := &http.Client{}
        req, _ := http.NewRequest("POST", "http://project:8000/api/me/new_game", bytes.NewBuffer(request))
        req.Header.Add("Authorization", "Bearer " + credentials.token)
        resp, err := client.Do(req)
        
        if err != nil {
            Respond(w, Message(false, "Post request error!"))
            return
        }
        
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)
        
        if err != nil {
            Respond(w, Message(false, "Json encode error!"))
            return
        }
        
        value := gjson.Get(string(body), "status").Bool()
        fmt.Println(value)
        
        if value {
            current_game.game_number = int(gjson.Get(string(body), "game.number").Int())
            fmt.Println("current_game.game_number ", current_game.game_number)
        }

        message := gjson.Get(string(body), "message").String()
        Respond(w, Message(false, message))
        
    }
}
