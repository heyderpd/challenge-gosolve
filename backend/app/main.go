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
	re := service.NewReader(os.Getenv("FILEPATH"))
	db := service.NewDatabase(re)
	r := router.SetupRouter(db)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
