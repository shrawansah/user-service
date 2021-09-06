package user

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"gojek.com/logger"

	"gojek.com/database"
	"gojek.com/repository"
)

type services interface {
	StartServing()
}

type simpleUserService struct {
	Database           *sql.DB
	UserRepository repository.UsersRepository
}

func NewSimpleUserService() services {

	defer logger.Info("SimpleUserService initialization complete")

	return simpleUserService{
		Database:           database.GetConnection(),
		UserRepository: repository.NewUsersRepository(database.GetConnection()),
	}
}

func (simpleUserService simpleUserService) StartServing() {

	logger.Info("Initializing to serve SimpleUserService")
	router := mux.NewRouter().StrictSlash(true)

	// user endpoints
	router.HandleFunc("/user/create", simpleUserService.CreateUserEndpointHandler).Methods("POST")

	logger.Info("Serving on port 8086")
	logger.Error(http.ListenAndServe(":8086", router))
}
