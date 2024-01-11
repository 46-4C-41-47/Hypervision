package main

import (
	"log"
	"net/http"
	"text/template"
)

const port = ":3000"

type User struct {
	username string
	password string
}

func Home(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./web/index.html")

	if err != nil {
		http.Error(response, err.Error(), 501)
	}

	tmpl.Execute(response, nil)
}

func Connect(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./web/user.html")

	if err != nil {
		http.Error(response, err.Error(), 501)
	}

	var user1 User
	user1.username = request.FormValue("username")
	user1.password = request.FormValue("password")

	log.Print("user : " + user1.username + " tried to connect with password : " + user1.password)

	tmpl.Execute(response, nil)
}

func main() {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/Connect", Connect)

	log.Print("Application started, go on http://localhost" + port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
