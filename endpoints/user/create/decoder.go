package createuser

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"gojek.com/logger"
)

func (createUserRequest *CreateUserRequest) Decode(r *http.Request) error {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := json.Unmarshal(reqBody, &createUserRequest); err != nil {
		logger.Error(err)
		return err
	}
	logger.Info("CreateUserRequest :: ", createUserRequest.ToString())
	return nil
}
