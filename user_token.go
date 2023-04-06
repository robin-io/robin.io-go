package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	urlx "net/url"
	"strings"
)

// The CreateUserToken function creates a usertoken with the meta data provided.

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

// sync UserToken (Update)

func (r *Robin) SyncUserToken(details UserToken) (UserTokenResponse, error) {
	body, err := json.Marshal(map[string]interface{}{
		"meta_data": details.MetaData,
	})

	if err != nil {
		return UserTokenResponse{}, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("PUT", fmt.Sprintf(`%s/chat/user_token/%s`, baseUrl, details.UserToken), bytes.NewBuffer(body))

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

func (r *Robin) UpdateDisplayPhoto(userToken string, photo string) (UserTokenResponse, error) {

	form := urlx.Values{}

	form.Add("display_photo", photo)

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf(`%s/chat/user_token/display_photo/%s`, baseUrl, userToken), strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("x-api-key", r.Secret)

	if err != nil {
		return UserTokenResponse{}, err
	}

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return UserTokenResponse{}, err
	}

	defer resp.Body.Close()

	bodyBuf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return UserTokenResponse{}, err
	}

	var newBody Response

	if err := json.Unmarshal(bodyBuf, &newBody); err != nil {
		return UserTokenResponse{}, err
	}

	if newBody.Error {
		return UserTokenResponse{}, errors.New(newBody.Msg)
	}

	return newBody.Data, nil

}

func (r *Robin) CheckUserTokenOnlineStatus(userTokens ...string) (map[string]string, error) {
	body, err := json.Marshal(map[string][]string{
		"user_tokens": userTokens,
	})

	if err != nil {
		return map[string]string{}, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/chat/user_token/get_online_status", baseUrl), bytes.NewBuffer(body))

	if err != nil {
		return map[string]string{}, err
	}

	req.Header.Set("x-api-key", r.Secret)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return map[string]string{}, err
	}

	defer resp.Body.Close()

	bodyBuf, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return map[string]string{}, err
	}

	var newBody UserTokenStatusResponse

	if err := json.Unmarshal(bodyBuf, &newBody); err != nil {
		return map[string]string{}, err
	}

	if newBody.Error {
		return map[string]string{}, errors.New(newBody.Msg)
	}

	return newBody.Data, nil
}
