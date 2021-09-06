package main

import (
	"gojek.com/logger"

	"gojek.com/service/user"
)

func main() {

	logger.Info("App Started")

	userService := user.NewUserService()
	userService.StartServing()

}
