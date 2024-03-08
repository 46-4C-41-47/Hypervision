package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"github.com/redis/go-redis/v9"
)

const port = "3000"

var ctx = context.Background()
var rdb *redis.Client

type User struct {
	username string `json:"username"`
	password string `json:"password"`
	role     string `json:"role"`
}

var currentUser User

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
		return
	}

	var u User
	u.username = request.FormValue("username")
	u.password = request.FormValue("password")

	log.Print("someone tried to connect as " + u.username)

	user, err := getUser(u.username)

	if err == redis.Nil || u.password != user.password {
		log.Print("connection to user " + u.username + " failed : username or password incorrect")

	} else if err != nil {
		panic(err)

	} else {
		log.Print("connection to " + u.username + " succeed")

		currentUser = user

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
}

func getUser(username string) (User, error) {
	var user User
	val, err := rdb.Get(ctx, username).Result()

	if err != nil {
		log.Print("The application was unable to retrieve users from Redis")
		return user, err
	}

	err = json.Unmarshal([]byte(val), &user)

	if err != nil {
		log.Print("The application was unable to parse the JSON retrieve from Redis")
		return user, err
	}

	return user, nil
}

func main() {
	initRedis()

	fs := http.FileServer(http.Dir("web"))
	http.Handle("/web/", http.StripPrefix("/web/", fs))

	http.HandleFunc("/", Home)
	http.HandleFunc("/Connect", Connect)

	log.Print("Application started, go on http://localhost:" + port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
