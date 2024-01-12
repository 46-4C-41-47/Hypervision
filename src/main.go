package main

import (
	"context"
	"log"
	"net/http"
	"text/template"

	"github.com/redis/go-redis/v9"
)

const port = ":3000"

var ctx = context.Background()
var rdb *redis.Client

type User struct {
	username string
	password string
}

func Home(response http.ResponseWriter, request *http.Request) {
	homeTemplate, err := template.ParseFiles("./web/index.html")

	if err != nil {
		http.Error(response, err.Error(), 501)
	}

	homeTemplate.Execute(response, nil)
}

func Connect(response http.ResponseWriter, request *http.Request) {
	tmpl, err := template.ParseFiles("./web/user.html")

	if err != nil {
		http.Error(response, err.Error(), 501)
	}

	var u User
	u.username = request.FormValue("username")
	u.password = request.FormValue("password")

	log.Print("user : " + u.username + " tried to connect with password : " + u.password)

	password, err := rdb.Get(ctx, u.username).Result()

	if err == redis.Nil || password != u.password {
		log.Print("username or password incorrect")

	} else if err != nil {
		panic(err)

	} else {
		log.Print("you are logged as " + u.username)
		tmpl.Execute(response, nil)
		return
	}

	Home(response, request)
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	err := rdb.Set(ctx, "user", "psswrd", 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	initRedis()

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
