package main

import (
	"fmt"
	"os"

	"challenge-gosolve/backend/app/config"
	"challenge-gosolve/backend/app/router"
	"challenge-gosolve/backend/app/service"
	"challenge-gosolve/backend/app/utils"
)

var db = make(map[string]string)

func main() {
	config.Init()
	utils.LoggerInit()
	re := service.NewReader(os.Getenv("FILEPATH"))
	db := service.NewDatabase(re)
	r := router.SetupRouter(db)
	r.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
