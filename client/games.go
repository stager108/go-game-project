
package main

import (
	"html/template"
	"log"
	"net/http"
    "io/ioutil"
    "github.com/tidwall/gjson"
	"strings"
    "bytes"
    "encoding/json"
	"fmt"
    "strconv"
)

var indexTemplate = template.Must(template.ParseFiles("index.tmpl"))

var gameTemplate = template.Must(template.Must(indexTemplate.Clone()).ParseFiles("game.tmpl"))

type Index struct {
	Title string
	Body  string
	Games []Game
	Links []Link
}

type Link struct {
	URL, Title string
}

type Game struct {
    Email string
    Number string
    Attempts string
    Score string
    Title string
}

var indexData = &Index{
    Title: "All your games are there!",
    Body:  "Hey, let's look at them.",
	}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    
    indexData.Games = indexData.Games[:0]
    indexData.Links = indexData.Links[:0]
    
    request, err := json.Marshal(map[string]string{
                "email":   credentials.email})
    
    if err != nil {
        Respond(w, Message(false, "Json decode error!"))
        return
    }
    
    client := &http.Client{}
    req, _ := http.NewRequest("POST", "http://project:8000/api/me/all_games", bytes.NewBuffer(request))
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
        
    count := int(gjson.Get(string(body), "data.#").Int())
    
    fmt.Println(count)
	
	for ind := 0; ind < count; ind++ {
		indexData.Links = append(indexData.Links, Link{
			URL:   "/my_games/my_game/" + strconv.Itoa(ind) ,
			Title: "GAME#" + strconv.Itoa(ind) ,
		})
        
        indexData.Games = append(indexData.Games, Game{
            Email:   gjson.Get(string(body), "data." + strconv.Itoa(ind) + ".email").String(),
            Number:   strconv.FormatInt(gjson.Get(string(body), "data." + strconv.Itoa(ind) + ".number").Int(), 10),
            Attempts:   strconv.FormatInt(gjson.Get(string(body), "data." + strconv.Itoa(ind) + ".attempts").Int(), 10),
            Score:   strconv.FormatInt(gjson.Get(string(body), "data." + strconv.Itoa(ind) + ".score").Int(), 10),
            Title: "GAME#" + strconv.Itoa(ind) ,
        })
	}
	
	if err := indexTemplate.Execute(w, indexData); err != nil {
		log.Println(err)
	}
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.URL.Path)
    splitted := strings.Split(r.URL.Path, "/")
	i, ok := strconv.Atoi(splitted[len(splitted) - 1])
    
    fmt.Println(i)
    
	if ok != nil {
		http.NotFound(w, r)
		return
	}
	nextgame := indexData.Games[i]
	if err := gameTemplate.Execute(w, nextgame); err != nil {
		log.Println(err)
	}
}
