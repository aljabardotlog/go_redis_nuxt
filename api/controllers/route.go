package controllers

import (
	"net/http"

	user "github.com/aljabardotlog/go_mongo_nuxt/controllers/user"
	"github.com/gorilla/mux"
)

func Server() {
	r := mux.NewRouter()

	s := r.PathPrefix("/api/v3").Subrouter()

	//private routes (authentication)
	////authorization

	//public routes
	s.HandleFunc("/user", user.GetAllData)
	s.HandleFunc("/user/{id}", user.GetSingleData)
	s.HandleFunc("/user", user.CreateData)
	s.HandleFunc("/user/{id}", user.PutData)
	s.HandleFunc("/user/{id}", user.DeleteData)

	http.ListenAndServe(":9001", r)

}
