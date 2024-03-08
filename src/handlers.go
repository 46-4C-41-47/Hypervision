package main

import (
	"html/template"
	"log"
	"net/http"
)

func Connect(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("username")
	password := r.FormValue("password")

	userTest := getUser(name)

	log.Print("someone tried to connect as " + name)

	if name == userTest.Name && password == userTest.Password {
		log.Print("connection succeed")
		loggedUser = userTest
		connectionFailed = false
		Home(w, r)
	} else {
		log.Print("connection failed")
		connectionFailed = true
		Login(w, r)
	}
}

func AddUser(w http.ResponseWriter, r *http.Request) {
	added_user := User{
		Name:     r.FormValue("username"),
		Password: r.FormValue("password"),
		Role:     r.FormValue("role"),
	}

	log.Print(loggedUser.Name + " tried to add an user named " + added_user.Name + " with role " + added_user.Role)

	err := rdb.Set(
		ctx,
		added_user.Name,
		`{"username":"`+added_user.Name+`","password":"`+added_user.Password+`","role":"`+added_user.Role+`"}`,
		0,
	).Err()

	if err != nil {
		log.Print("The application was unable to add user " + added_user.Name + " to Redis")
	} else {
		log.Print("User " + added_user.Name + " successfully added to Redis")
	}
}

func Dashboard(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./front/dashboard.html"))
	tmpl.Execute(w, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	if loggedUser.Name == "" {
		return
	}

	tmpl := template.Must(template.ParseFiles("./front/home.html"))
	tmpl.Execute(w, map[string]bool{"IsAdmin": loggedUser.Role == Admin || loggedUser.Role == SuperAdmin})
}

func Login(w http.ResponseWriter, r *http.Request) {
	if loggedUser.Name == "" {
		tmpl := template.Must(template.ParseFiles("./front/login.html"))
		tmpl.Execute(w, map[string]bool{"Failed": connectionFailed})
	} else {
		Home(w, r)
	}
}

func Users(w http.ResponseWriter, r *http.Request) {
	if loggedUser.Role == Visitor {
		Home(w, r)
		return
	}

	a := UserPageData{
		getAllUsers(),
		loggedUser.Role == Admin || loggedUser.Role == SuperAdmin,
		loggedUser.Role == SuperAdmin,
	}

	tmpl := template.Must(template.ParseFiles("./front/users.html"))
	tmpl.Execute(w, map[string]UserPageData{"Data": a})
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if loggedUser.Role == SuperAdmin {
		removeUser(r.FormValue("username"))
	}
}

func SignOut(w http.ResponseWriter, r *http.Request) {
	log.Print(loggedUser.Name + " logged out")
	loggedUser = User{}
	Login(w, r)
}
