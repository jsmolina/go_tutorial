package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)

type Article struct {
    Id      string `json:"id"`
    Title string `json:"title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type ErrorResponse struct {
    Error string `json:"error"`
}

// let's declare a global Articles array
var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    var result interface{}
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            result = &article
        }
    }

    if (result == nil) {
        w.WriteHeader(404)
        var error = ErrorResponse{Error: "Not found"}
        result = &error
    }

    fmt.Fprint(w, json.NewEncoder(w).Encode(result))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}


func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)

    log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")
    Articles = [] Article {
        Article{Id: "1",Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2",Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}
