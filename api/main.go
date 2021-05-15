package main

import (
	"fmt"

	"github.com/aljabardotlog/go_mongo_nuxt/controllers"
)

func main() {
	fmt.Println("Running in server :9001")

	controllers.Server()
}
