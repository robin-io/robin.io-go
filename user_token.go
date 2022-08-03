package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
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

func (r *Robin) UpdateDisplayPhoto(userToken string, photo *multipart.FileHeader) (UserTokenResponse, error) {

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("display_photo", photo.Filename)

	if err != nil {
		return UserTokenResponse{}, err
	}

	file, err := photo.Open()

	if err != nil {
		return UserTokenResponse{}, err
	}

	_, err = io.Copy(part, file)

	if err != nil {
		return UserTokenResponse{}, err
	}

	err = writer.Close()

	if err != nil {
		return UserTokenResponse{}, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf(`%s/chat/user_token/display_photo/%s`, baseUrl, userToken), body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
