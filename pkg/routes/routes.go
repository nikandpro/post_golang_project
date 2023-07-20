package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"backend/basic/pkg/handlers"
)

func InitRoutes() http.Handler {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/", handlers.Index).Methods("GET")
	rtr.HandleFunc("/create", handlers.Create).Methods("GET")
	rtr.HandleFunc("/save_article", handlers.Save_article).Methods("POST")
	rtr.HandleFunc("/post/{id:[0-9]+}", handlers.Show_post).Methods("GET")
	rtr.HandleFunc("/sign", handlers.Sign).Methods("GET")
	// rtr.HandleFunc("/user/{id:[0-9]+}", handlers.IndexUser).Methods("GET")
	rtr.HandleFunc("/sign_user",handlers.Sign_user).Methods("POST")

	//authentication
	rtr.HandleFunc("/login", handlers.Login).Methods("POST")

	//CRUD User
	rtr.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	rtr.HandleFunc("/users/{id}",handlers.GetUser).Methods("GET")
	rtr.HandleFunc("/users" , handlers.CreateUser).Methods("POST")
	rtr.HandleFunc("/users/{id}",handlers.UpdateUser).Methods("PUT")
	rtr.HandleFunc("/users/{id}",handlers.DeleteUser).Methods("DELETE")


	return rtr
}