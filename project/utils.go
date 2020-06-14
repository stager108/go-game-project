package main

import (
	"encoding/json"
	"net/http"
    "text/template"
)

type page struct {
    Title     string
    Message string
}

func Message(status bool, message string) (map[string]interface{}) {
	return map[string]interface{} {"status" : status, "message" : message}
}

func Respond(w http.ResponseWriter, data map[string] interface{})  {
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}


func DisplayMessage(w http.ResponseWriter, message string){
    
        
    w.Header().Add("Content Type", "text/html")

    templates := template.New("template")

    page := page{Title: "Message", Message: message}
    templates.Execute(w, page)
        

   // http.ListenAndServe(":8000", nil)
}
