package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/novalagung/gubrak/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var validate *validator.Validate

type App struct {
	DB *gorm.DB
}

func (a *App) Initialize() {
	dsn := "root:LinuxKU13!@tcp(docker.for.mac.localhost:3306)/intranet_v2?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = db
	validate = validator.New()
}

func (app *App) RegisterCustomValidation() {
	validate.RegisterValidation("username", ValidateUsername)
}

func (app *App) ValidationMiddleware(request interface{}, messages map[string]string, next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var errorResponse ErrorResponseItem

		err := ParseRequest(w, r, request)
		if err != nil {
			errorResponse = ErrorResponseItem{
				Code:   400,
				Status: "Error 400",
			}
			fmt.Println(err)
			SendErrorResponse(w, errorResponse, 400)
			return
		}

		err = validate.Struct(request)
		if err != nil {
			if _, ok := err.(*validator.InvalidValidationError); ok {
				log.Printf("%v\n", err)
				errorResponse = ErrorResponseItem{
					Code:   403,
					Status: "Bad Request: Just bad request",
				}
			} else if len(err.(validator.ValidationErrors)) > 0 {
				tag := err.(validator.ValidationErrors)[0].Tag()
				field := err.(validator.ValidationErrors)[0].Field()
				messageKey := fmt.Sprintf("%s_%s", field, tag)
				message := gubrak.From(messages).Filter(func(value string, key string) bool {
					return key == messageKey
				}).Result()

				var messageStr string

				messageMap, _ := message.(map[string]string)
				if len(messageMap) == 0 {
					messageStr = fmt.Sprintf("error message not found for field: %s", field)
				} else {
					messageStr = messageMap[messageKey]
				}

				errorResponse = ErrorResponseItem{
					Code:   422,
					Status: messageStr,
				}
			}

			SendErrorResponse(w, errorResponse, errorResponse.Code)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}
