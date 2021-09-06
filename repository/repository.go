package repository

import (
	db "gojek.com/database"
)

type repository struct {
	UsersRepository        UsersRepository
}

var Repositories = repository{
	UsersRepository:        NewUsersRepository(db.GetConnection()),
}
