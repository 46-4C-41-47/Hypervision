package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name     string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserPageData struct {
	Users        []User
	IsAdmin      bool
	IsSuperadmin bool
}

const port = "3000"
const Visitor = "visitor"
const Admin = "admin"
const SuperAdmin = "superadmin"

var ctx = context.Background()
var rdb *redis.Client
var loggedUser = User{}
var connectionFailed = false

func getAllUsers() []User {
	v, err := rdb.Keys(ctx, "*").Result()
	var users []User

	if err != nil {
		log.Print("The application was unable to retrieve keys from Redis")
		return nil
	}

	for _, element := range v {
		users = append(users, getUser(element))
	}

	return users
}

func getUser(username string) User {
	v, redisErr := rdb.Get(ctx, username).Result()
	user := User{}

	if redisErr != nil {
		log.Print("The application was unable to retrieve users from Redis")
		return user
	}

	jsonErr := json.Unmarshal([]byte(v), &user)

	if jsonErr != nil {
		log.Print("The application was unable to parse JSON from Redis")
	}

	return user
}

func removeUser(username string) {
	_, redisErr := rdb.Del(ctx, username).Result()

	if redisErr != nil {
		log.Print("The application was unable to remove user " + username + " from Redis")
	} else {
		log.Print(username + " was removed from user list")
	}
}

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})
}

func main() {
	initRedis()

	getAllUsers()

	http.HandleFunc("/", Login)
	http.HandleFunc("/Users", Users)
	http.HandleFunc("/AddUser", AddUser)
	http.HandleFunc("/DeleteUser", DeleteUser)
	http.HandleFunc("/Dashboard", DeleteUser)
	http.HandleFunc("/Connect", Connect)
	http.HandleFunc("/Home", Home)
	http.HandleFunc("/SignOut", SignOut)

	log.Print("application started on : http://localhost:" + port + "\n")

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
