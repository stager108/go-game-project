package main

import (
	"net/http"
    "html/template"
	"encoding/json"
    "fmt"
    "strconv"
    "io/ioutil"
    "bytes"
    "github.com/tidwall/gjson"
)


var GetNewAttempt = func(w http.ResponseWriter, r *http.Request) {
    
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        t, _ := template.ParseFiles("new_attempt.gtpl")
        t.Execute(w, nil)
    } else {
        r.ParseForm()

        request, err := json.Marshal(map[string]string{
                    "email":   credentials.email,
                    "game":    strconv.Itoa(current_game.game_number),
        })
        
        if err != nil {
            Respond(w, Message(false, "Json decode error!"))
            return
        }
        
        client := &http.Client{}
        req, _ := http.NewRequest("POST", "http://project:8000/api/me/new_attempt", bytes.NewBuffer(request))
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
            current_game.current_attempt = int(gjson.Get(string(body), "attempt.number").Int())
            current_game.Current_word = gjson.Get(string(body), "attempt.word").String()
        }

        message := gjson.Get(string(body), "message").String()
        Respond(w, Message(true, message))
    }
}


var SendAnswer = func(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //get request method
    if r.Method == "GET" {
        fmt.Println("method:", current_game.Current_word)
        fmt.Println("method:", current_game.current_attempt)
        t, _ := template.ParseFiles("submit.gtpl")
        
        if err := t.Execute(w, current_game); err != nil {
            Respond(w, Message(false, "Template error!"))
            return
        }
    } else {
        r.ParseForm()

        request, err := json.Marshal(map[string]string{
                    "email":   credentials.email,
                    "game": strconv.Itoa(current_game.game_number),
                    "number": strconv.Itoa(current_game.current_attempt),
                    "word": r.FormValue("answer"),
        })
        
        if err != nil {
            Respond(w, Message(false, "Json decode error!"))
            return
        }
        
        client := &http.Client{}
        req, _ := http.NewRequest("POST", "http://project:8000/api/me/send_answer", bytes.NewBuffer(request))
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
        message := gjson.Get(string(body), "message").String()
        Respond(w, Message(true, message))
    }
}





