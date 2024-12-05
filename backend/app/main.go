package main

import (
	"fmt"
	"os"

	"challenge-gosolve/backend/app/config"
	"challenge-gosolve/backend/app/router"
	"challenge-gosolve/backend/app/service"
)

var db = make(map[string]string)

func main() {
	config.Init()
	r := router.SetupRouter(service.NewDatabase())
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
