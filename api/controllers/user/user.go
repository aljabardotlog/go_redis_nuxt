package user

import (
	"net/http"

	rt "github.com/aljabardotlog/go_mongo_nuxt/handlers"
	"github.com/aljabardotlog/go_mongo_nuxt/models"
)

func (user *User) CreateData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := UserResponse{
		ID:       1,
		Username: "ini create",
		Password: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) GetAllData(w http.ResponseWriter, r *http.Request) {

	//logic
	DataRaw := models.User{}
	if err := user.App.DB.Find(&DataRaw).Error; err != nil {
		errorResponse := rt.ErrorResponseItem{
			Code:   404,
			Status: "Not Found",
		}
		rt.SendErrorResponse(w, errorResponse, http.StatusBadRequest)
		return
	}

	payload := UserResponse{
		ID:       DataRaw.ID,
		Username: DataRaw.Username,
		Password: DataRaw.Password,
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) GetSingleData(w http.ResponseWriter, r *http.Request) {
	//logic

	payload := UserResponse{
		ID:       1,
		Username: "ini get single",
		Password: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) DeleteData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := UserResponse{
		ID:       1,
		Username: "ini delete",
		Password: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) PutData(w http.ResponseWriter, r *http.Request) {

	var request UserRequest

	err := rt.ParseRequest(w, r, &request)
	if err != nil {
		errorResponse := rt.ErrorResponseItem{
			Code:   404,
			Status: "Not Found",
		}
		rt.SendErrorResponse(w, errorResponse, 400)
	}

	payload := UserResponse{
		ID:       0,
		Username: request.Username,
		Password: request.Password,
	}

	rt.SendResponse(w, payload, http.StatusOK)
}
