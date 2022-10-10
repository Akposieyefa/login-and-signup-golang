package routes

import (
	"log"
	"net/http"

	"github.com/akposieyefa/login-and-signup/controllers"
	"github.com/akposieyefa/login-and-signup/middleware"
	"github.com/gorilla/mux"
)

func LoadRouters() {

	r := mux.NewRouter()
	r.Use(middleware.Middleware)

	r.HandleFunc("/create-account", controllers.CreateUserAccount).Methods("POST")
	r.HandleFunc("/login", controllers.LoginUser).Methods("POST")

	subRouter := r.PathPrefix("/sub_router/").Subrouter()
	subRouter.Use(middleware.JwtVerify)
	subRouter.HandleFunc("/user", controllers.FetchUsers).Methods("GET")
	r.HandleFunc("/get-users", controllers.LoginUser).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:9000", r))

}
