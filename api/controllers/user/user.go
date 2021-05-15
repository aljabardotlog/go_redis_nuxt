package user

import (
	"net/http"

	rt "github.com/aljabardotlog/go_mongo_nuxt/handlers"
)

func CreateData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := rt.Response{
		Data: userResponse{
			ID:        "1",
			FirstName: "Andi",
			LasttName: "Aljabar",
		},
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func GetAllData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := rt.Response{
		Data: userResponse{
			ID:        "1",
			FirstName: "Andi",
			LasttName: "Aljabar",
		},
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func GetSingleData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := rt.Response{
		Data: userResponse{
			ID:        "1",
			FirstName: "Andi",
			LasttName: "Aljabar",
		},
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func DeleteData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := rt.Response{
		Data: userResponse{
			ID:        "1",
			FirstName: "Andi",
			LasttName: "Aljabar",
		},
	}

	rt.SendResponse(w, payload, http.StatusOK)
}

func PutData(w http.ResponseWriter, r *http.Request) {

	//logic

	payload := rt.Response{
		Data: userResponse{
			ID:        "1",
			FirstName: "Andi",
			LasttName: "Aljabar",
		},
	}

	rt.SendResponse(w, payload, http.StatusOK)
}
