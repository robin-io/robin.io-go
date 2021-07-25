package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *Robin) CreateUserToken(details UserToken) (UserTokenResponse, error) {
	body, err := json.Marshal(map[string]interface{}{
		"meta_data": details.MetaData,
	})

	if err != nil {
		return UserTokenResponse{}, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(`%s/chat/user_token`, baseUrl), bytes.NewBuffer(body))

	if err != nil {
		return UserTokenResponse{}, err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return UserTokenResponse{}, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return UserTokenResponse{}, err
	}

	var newBody Response

	if err := json.Unmarshal(body, &newBody); err != nil {
		return UserTokenResponse{}, err
	}

	if newBody.Error {
		return UserTokenResponse{}, errors.New(newBody.Msg)
	}

	return newBody.Data, nil

}

// get conversations (Get User Token)

func (r *Robin) GetUserToken(userToken string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(`%s/chat/user_token/%s`, baseUrl, userToken), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newBody Response

	if err := json.Unmarshal(body, &newBody); err != nil {
		return nil, err
	}

	if newBody.Error {
		return nil, errors.New(newBody.Msg)
	}

	return newBody.Data, nil
}