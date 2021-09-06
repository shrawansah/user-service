package createuser

import (
	"regexp"
	"strings"
)

func (createUserRequest *CreateUserRequest) Validate() []string {
	var err []string

	if strings.TrimSpace(createUserRequest.Name) == "" {
		err = append(err, "name can not be empty")
	}
	if createUserRequest.Email == "" {
		err = append(err, "email required")

	} else {
		rxEmail := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		if !rxEmail.MatchString(createUserRequest.Email) {
			err = append(err, "Invalid email address")
		}
	}

	return err
}