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

type userService struct {
	Database       *sql.DB
	UserRepository repository.UsersRepository
}

func NewUserService() services {

	defer logger.Info("SimpleUserService initialization complete")

	return userService{
		Database:       database.GetConnection(),
		UserRepository: repository.NewUsersRepository(database.GetConnection()),
	}
}

func (userService userService) StartServing() {

	logger.Info("Initializing to serve SimpleUserService")
	router := mux.NewRouter().StrictSlash(true)

	// user endpoints
	router.HandleFunc("/user/create", userService.CreateUserEndpointHandler).Methods("POST")

	logger.Info("Serving on port 8086")
	logger.Error(http.ListenAndServe(":8086", router))
}
