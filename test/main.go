package main

import (
	"net/http"
	"encoding/json"
    "fmt"
    "io/ioutil"
    "bytes"
    "github.com/tidwall/gjson"
    "strconv"
)

type ClientCredentials struct{
    email string
    token string
} 

type CurrentGame struct{
    game_number string
    current_attempt string
    current_word string
} 

var credentials = &ClientCredentials{}

var current_game = &CurrentGame{}


var hostname = "project"

func AuthenticateTestGoodParameters(test_email, test_password, api_path  string ) {
    
    fmt.Println("test: "+ api_path + " with good parameters") //get request method

    request, err := json.Marshal(map[string]string{
                "email":   test_email,
                "password" : test_password})
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
        return
    }

    resp, err := http.Post("http://" + hostname + ":8000/api/user" + api_path,
        "application/json", bytes.NewBuffer(request))
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
        
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
        return
    }
        
    value := gjson.Get(string(body), "status").Bool()
    message := gjson.Get(string(body), "message").String()
    fmt.Println("server message: " + message)
    
    if value {
        email := gjson.Get(string(body), "account.email").String()
        if email != test_email {
            fmt.Println("test failed: wrong email!")
            fmt.Println()
            return
        }
        
        credentials.email = gjson.Get(string(body), "account.email").String()
        credentials.token = gjson.Get(string(body), "account.token").String()
        fmt.Println("test passed!")
    } else {
        fmt.Println("test failed: wrong status!")
    }

    fmt.Println()
    
}


func AuthenticateTestBadParameters(test_email, test_password, api_path string) {
    
    fmt.Println("test: "+ api_path + " with bad parameters")

    request, err := json.Marshal(map[string]string{
                "email":   test_email,
                "password" : test_password})
    
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
        return
    }

    resp, err := http.Post("http://" + hostname + ":8000/api/user" + api_path,
        "application/json", bytes.NewBuffer(request))
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
        
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
        return
    }
        
    value := gjson.Get(string(body), "status").Bool()
    message := gjson.Get(string(body), "message").String()
    fmt.Println("server message: " + message)
        
    if !value {
        fmt.Println("test passed!")
    }  else {
        fmt.Println("test failed: wrong status!")
    }
    fmt.Println()
}


func NewGameTest(test_email string){
    
    fmt.Println("test: /new_game with good parameters")
    
    request, err := json.Marshal(map[string]string{
                    "email":  test_email})
        
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
    }
        
    client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://" + hostname + ":8000/api/me/new_game", bytes.NewBuffer(request))
    req.Header.Add("Authorization", "Bearer " + credentials.token)
    resp, err := client.Do(req)
    
    if err != nil {
        fmt.Println("test failed: post request error!")
        fmt.Println()
        return
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    
    if err != nil {
        fmt.Println("test failed: error while reading response body!")
        fmt.Println()
        return
    }
    
    message := gjson.Get(string(body), "message").String()
    fmt.Println("server message: " + message)
    value := gjson.Get(string(body), "status").Bool()

    if value {
        current_game.game_number = strconv.Itoa(int(gjson.Get(string(body), "game.number").Int()))
        fmt.Println("test passed!")
    } else {
        fmt.Println("test failed: wrong status!")
    }
    fmt.Println()
}

func NewAttemptTest(test_email, game_number string, good_parameters bool) {
    
    if good_parameters {
        fmt.Println("test: /play with good parameters")
    } else {
        fmt.Println("test: /play with bad parameters")
    }
    
    request, err := json.Marshal(map[string]string{
                "email":   test_email,
                "game":    game_number,
    })
    
    if err != nil {
        fmt.Println("test failed: json encode error!")
        fmt.Println()
        return
    }

    client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://" + hostname + ":8000/api/me/new_attempt", bytes.NewBuffer(request))
    req.Header.Add("Authorization", "Bearer " + credentials.token)
    resp, err := client.Do(req)
        
    if err != nil {
        fmt.Println("test failed: post request error!")
        fmt.Println()
        return
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
        
    if err != nil {
        fmt.Println("test failed: error while reading response body!")
        fmt.Println()
        return
    }
    
    value := gjson.Get(string(body), "status").Bool()
    
    message := gjson.Get(string(body), "message").String()
    fmt.Println("server message: " + message)
    
    if good_parameters {
        if value {
            current_game.current_attempt = strconv.Itoa(int(gjson.Get(string(body), "attempt.number").Int()))
            current_game.current_word = gjson.Get(string(body), "attempt.word").String()
        } else {
            fmt.Println("test failed: wront status!")
            fmt.Println()
            return
        }
    } else {
        if !value {
            fmt.Println("test failed: wront status!")
            fmt.Println()
            return
        }
    }

    fmt.Println("test passed!")
    fmt.Println()
    
    if good_parameters {
        SendTestGoodParameters(test_email, game_number, current_game.current_attempt, current_game.current_word)
    }
}


func SendTestGoodParameters(test_email, game_number, attempt, word string ) {
fmt.Println("test: /send with good parameters")
    
    request, err := json.Marshal(map[string]string{
                "email":   test_email,
                "game": game_number,
                "number": attempt,
                "word": word,
    })
    
    if err != nil {
        fmt.Println("test failed: post request error!")
        return
    }
    
    client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://" + hostname + ":8000/api/me/send_answer", bytes.NewBuffer(request))
    req.Header.Add("Authorization", "Bearer " + credentials.token)
    resp, err := client.Do(req)
    
    if err != nil {
        fmt.Println("Post request error!")
        fmt.Println()
        return
    }
    
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    message := gjson.Get(string(body), "message").String()
    fmt.Println("server message: " + message)
    
    if err != nil {
        fmt.Println("Json encode error!")
        fmt.Println()
        return
    }
    
    value := gjson.Get(string(body), "status").Bool()
    
    if value {
        fmt.Println("test passed!")
    }  else {
        fmt.Println("test failed: wrong status!")
    }

    fmt.Println()
    
}

func test1(){
    
    AuthenticateTestGoodParameters("xx@xx", "123456", "/new")
    AuthenticateTestBadParameters("aaaa", "123456", "/new")
    AuthenticateTestBadParameters("aa@aa", "12345", "/new")
    AuthenticateTestBadParameters("aa@a@a", "123456", "/new")
    AuthenticateTestBadParameters("xx@xx", "123456", "/new")
}

func test2(){

    AuthenticateTestBadParameters("aa1@aa", "123456", "/login") 
    AuthenticateTestGoodParameters("aa1@aa", "123456", "/new")
    AuthenticateTestGoodParameters("bb@bb", "654321", "/new")
    AuthenticateTestGoodParameters("aa1@aa", "123456", "/login") 
    AuthenticateTestGoodParameters("bb@bb", "654321", "/login") 
    AuthenticateTestBadParameters("bb@bb", "123456", "/login") 
}

func test3(){
    AuthenticateTestGoodParameters("xx@xx", "123456", "/login")
    NewGameTest("xx@xx")
}


func test4(){
    AuthenticateTestGoodParameters("xx@xx", "123456", "/login")
    NewAttemptTest("xx@xx", "100500", false)
}

func test5(){
    AuthenticateTestGoodParameters("xx@xx", "123456", "/login")
    NewAttemptTest("xx@xx", current_game.game_number, true)
}


func main() {
    test1()
    test2()
    test3()
    test4()
    test5()
}
 
