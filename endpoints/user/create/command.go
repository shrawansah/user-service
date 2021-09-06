package createuser

import (
	"encoding/json"
	"strings"

	"gojek.com/logger"
)

type CreateUserCommand struct {
	Name 				string      `json:"name"`
	Email			 	string		`json:"email"`
}

func (createUserCommand *CreateUserCommand) ToString() string {
	bytes, _ := json.Marshal(createUserCommand)
	return string(bytes)
}


func (command *CreateUserCommand) BuildFromRequest(request *CreateUserRequest) {

	command.Name = strings.TrimSpace(request.Name)
	command.Email = strings.TrimSpace(request.Email)

	logger.Info("CreateUserCommand :: ", command.ToString())
}