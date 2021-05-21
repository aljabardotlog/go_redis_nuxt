package user

import (
	handlers "github.com/aljabardotlog/go_mongo_nuxt/handlers"
)

type User struct {
	App *handlers.App
}

type userResponse struct {
	ID        uint   `json:"_id"`
	FirstName string `json:"firstname"`
	LasttName string `json:"lastname"`
}
