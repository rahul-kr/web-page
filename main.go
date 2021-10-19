package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Param struct {
	UrlParam string
}

type postRequestData struct {
	ArticleId string `json:"article_id"`
}

var router = mux.NewRouter()

func main() {

	router.HandleFunc("/", homePage)
	router.HandleFunc("/page/{key}", PageHandler)
	log.Println("about to start the application...")
	log.Fatal(http.ListenAndServe(":8070", router))
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	urlParam := params["key"]

	P := Param{UrlParam: urlParam}
	custTemplate, err := template.ParseFiles("webPage.html")
	if err != nil {
		log.Println("Error in file parsing")
	}
	err = custTemplate.Execute(w, P)
	if err != nil {
		log.Println(err)

	}
	fmt.Println("Performing Http Post...")
	todo := postRequestData{urlParam}
	jsonReq, err := json.Marshal(todo)
	resp, err := http.Post("http://localhost:8888/counter/v1/statistics", "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love Golang")
}
