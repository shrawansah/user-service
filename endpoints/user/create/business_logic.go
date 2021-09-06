package createuser

import (
	"net/http"

	"gojek.com/errors"
	"gojek.com/logger"
	. "gojek.com/repository"
	"gojek.com/repository/models"
)

func (command *CreateUserCommand) ExecuteBusinessLogic() (*models.User, errors.BusinessLogicError) {
	
	user := models.User {
		Name: command.Name,
		EmailID: command.Email,
	}
	businessError := errors.BusinessLogicError{}
	defer func() {
		if !businessError.IsNil() {
			logger.Info("BusinessLogic error :: ", businessError)
		}
	}()

	users, err := Repositories.UsersRepository.GetUsers("email_id = ? ", command.Email)
	if err != nil {
		logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &user, businessError
	}
	if len(users) > 0 {
		businessError.ClientHTTPCode = http.StatusBadRequest
		businessError.ClientMessage = "user with same email already exists"

		return &user, businessError
	}

	if err := Repositories.UsersRepository.PutUser(&user, nil); err != nil {
		logger.Error(err)
		businessError.ClientHTTPCode = http.StatusInternalServerError
		businessError.ClientMessage = "I am a teacup!"

		return &user, businessError
	}

	return &user, businessError
}