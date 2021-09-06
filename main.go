package main

import (
	"gojek.com/logger"

	"gojek.com/service/user"
)

func main() {

	logger.Info("App Started")

	simplePaylaterService := user.NewSimpleUserService()
	simplePaylaterService.StartServing()

}
