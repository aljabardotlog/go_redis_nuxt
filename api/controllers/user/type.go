package user

import (
	handlers "github.com/aljabardotlog/go_mongo_nuxt/handlers"
)

type User struct {
	App *handlers.App
}

var ErrorMessages map[string]string = map[string]string{
	"Username_required": "Username belum di-isi",
}

type UserResponse struct {
	ID       uint   `json:"_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password"`
}
