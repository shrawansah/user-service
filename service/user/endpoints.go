package user

import (
	"encoding/json"
	"net/http"

	createuser "gojek.com/endpoints/user/create"

	"gojek.com/logger"
)

// CreateUserEndpointHandler
/**
User Endpoints Begin
**/
func (service userService) CreateUserEndpointHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var createUserRequest createuser.CreateUserRequest
	var response interface{}

	defer func() {
		bytes, _ := json.Marshal(response)
		logger.Info(string(bytes))
		json.NewEncoder(w).Encode(response)
	}()

	// decode
	if err := createUserRequest.Decode(r); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response = "invalid request format"
		return
	}

	// validate
	if errors := createUserRequest.Validate(); len(errors) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		response = errors
		return
	}

	// command
	createUserCommand := createuser.CreateUserCommand{}
	createUserCommand.BuildFromRequest(&createUserRequest)

	// business logic
	response, businessError := createUserCommand.ExecuteBusinessLogic()
	if !businessError.IsNil() {
		w.WriteHeader(businessError.ClientHTTPCode)
		response = businessError.ClientMessage
		return
	}

	w.WriteHeader(http.StatusCreated)
}

/**
User Endpoints Ends
**/
