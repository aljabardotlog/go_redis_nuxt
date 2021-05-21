package user

import (
	"net/http"

	rt "github.com/aljabardotlog/go_mongo_nuxt/handlers"
	"github.com/aljabardotlog/go_mongo_nuxt/models"
)

func (user *User) CreateData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := userResponse{
		ID:        1,
		FirstName: "Andi",
		LasttName: "Aljabar",
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

	payload := userResponse{
		ID:        DataRaw.ID,
		FirstName: DataRaw.Username,
		LasttName: DataRaw.Password,
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) GetSingleData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := userResponse{
		ID:        1,
		FirstName: "Andi",
		LasttName: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) DeleteData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := userResponse{
		ID:        1,
		FirstName: "Andi",
		LasttName: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func (user *User) PutData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := userResponse{
		ID:        1,
		FirstName: "Andi",
		LasttName: "Aljabar",
	}

	rt.SendResponse(w, payload, http.StatusOK)
}
