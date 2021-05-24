package controllers

import (
	"log"
	"net/http"

	"github.com/aljabardotlog/go_mongo_nuxt/controllers/user"
	"github.com/aljabardotlog/go_mongo_nuxt/handlers"
	hnd "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Server() {

	r := mux.NewRouter()
	s := r.PathPrefix("/api/v3").Subrouter()

	app := handlers.App{}
	app.Initialize()

	User := user.User{App: &app}
	//private routes (authentication)

	////authorization

	//public routes
	s.HandleFunc("/user", User.CreateData).Methods("POST")
	s.HandleFunc("/user", User.GetAllData).Methods("GET")
	s.HandleFunc("/user/{id}", User.GetSingleData).Methods("GET")
	s.HandleFunc(
		"/user/{id}",
		app.ValidationMiddleware(
			&user.UserRequest{},
			user.ErrorMessages,
			http.HandlerFunc(User.PutData),
		),
	).Methods("PUT")
	s.HandleFunc("/user/{id}", User.DeleteData).Methods("DELETE")

	headers := hnd.AllowedHeaders([]string{"Content-Type", "Authorization"})
	// methods := hnd.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "DELETE", "OPTIONS"})
	methods := hnd.AllowedMethods([]string{"POST"})
	origins := hnd.AllowedOrigins([]string{"*"})

	log.Fatal(http.ListenAndServe(":9001", hnd.CORS(methods, origins, headers)(r)))

}
