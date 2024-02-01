package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
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
var users []User

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

	loadUsers()

	var u User
	u.username = request.FormValue("username")
	u.password = request.FormValue("password")

	log.Print("someone tried to connect as " + u.username)

	user := getUser(u.username, u.password)

	if err == redis.Nil || user.username == "noUser" {
		log.Print("connection to user " + u.username + " failed : username or password incorrect")

	} else if err != nil {
		panic(err)

	} else {
		log.Print("connection to " + u.username + " succeed")

		currentUser.username = u.username

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

func loadUsers() {
	val, redisErr := rdb.Get(ctx, "users").Result()

	if redisErr != nil {
		log.Print("The application was unable to retrieve users from Redis")
		return
	}

	err := json.Unmarshal([]byte(val), &users)

	if err != nil {
		log.Print("The application was unable to parse the JSON retrieve from Redis")
		return
	}

	for i := 0; i < len(users); i++ {
		log.Print("index : " + strconv.Itoa(i) + " " + users[i].username)
	}
}

func getUser(username string, password string) User {
	for i := 0; i < len(users); i++ {
		if users[i].username == username && users[i].password == password {
			return users[i]
		}
	}

	var noUser User
	noUser.username = "noUser"
	return noUser
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
