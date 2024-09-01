package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"someweb/cmd/handlers"

	"github.com/joho/godotenv"
)

func main() {
	loadEnvVariables()
	routeHandler()
	runApp()
}

func loadEnvVariables() {
	//godotenv simple lib to read from .env
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(os.Getenv("PORT"))

	//set con string
	handlers.SetDBConn(os.Getenv("DB"))
}

func routeHandler() {
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("../views/assets"))))
	http.HandleFunc("/", routeIndex)
	http.HandleFunc("/posts", routeGetArticles)
}

func routeIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../views/index.html"))
	err := tmpl.Execute(w, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func routeGetArticles(w http.ResponseWriter, r *http.Request) {
	posts := handlers.GetArticles()

	tmpl := template.Must(template.ParseFiles("../views/index.html"))
	err := tmpl.Execute(w, posts)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func runApp() {
	myport := os.Getenv("PORT")
	fmt.Printf("server start at localhost%s", myport)
	err := http.ListenAndServe(myport, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}
