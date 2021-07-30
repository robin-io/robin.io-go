package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *Robin) CreateConversation(senderName, senderToken, receiverToken, receiverName string) (interface{}, error) {
	body, err := json.Marshal(map[string]string{
		"sender_name":   senderName,
		"sender_token":     senderToken,
		"receiver_token":   receiverToken,
		"receiver_name": receiverName,
	})

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(`%s/conversation`, baseUrl), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newBody ConversationResponse

	if err := json.Unmarshal(body, &newBody); err != nil {
		return nil, err
	}

	if newBody.Error {
		return nil, errors.New(newBody.Msg)
	}

	return newBody.ConversationData, nil
}

func (r *Robin) GetConversationMessages(id string) ([]MessageResponseData, error) {
	if len(id) == 0 {
		return []MessageResponseData{}, errors.New("id cannot be empty")
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(`%s/conversation/messages/%s`, baseUrl, id), nil)

	if err != nil {
		return []MessageResponseData{}, err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return []MessageResponseData{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []MessageResponseData{}, err
	}

	var newBody MessageResponse

	if err := json.Unmarshal(body, &newBody); err != nil {
		return []MessageResponseData{}, err
	}

	if newBody.Error {
		return []MessageResponseData{}, errors.New(newBody.Msg)
	}

	return newBody.MessageData, nil
}

func (r *Robin) SearchConversation(id, text string) ([]MessageResponseData, error) {

	/*
		an empty text string returns all messages in the conversation
	*/

	if len(id) == 0 {
		return []MessageResponseData{}, errors.New("id cannot be empty")
	}

	body, err := json.Marshal(map[string]string{
		"text": text,
	})

	if err != nil {
		return []MessageResponseData{}, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(`%s/chat/search/message/%s`, baseUrl, id), bytes.NewBuffer(body))

	if err != nil {
		return []MessageResponseData{}, err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return []MessageResponseData{}, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return []MessageResponseData{}, err
	}

	var newBody MessageResponse

	if err := json.Unmarshal(body, &newBody); err != nil {
		return []MessageResponseData{}, err
	}

	if newBody.Error {
		return []MessageResponseData{}, errors.New(newBody.Msg)
	}

	return newBody.MessageData, nil
}

func (r *Robin) DeleteMessage(id string) error {
	if len(id) == 0 {
		return errors.New("id cannot be empty")
	}

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", fmt.Sprintf(`%s/chat/message/%s`, baseUrl, id), nil)

	if err != nil {
		return err
	}

	req.Header.Set("x-api-key", r.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return err
	}

	var newBody MessageResponse

	if err := json.Unmarshal(body, &newBody); err != nil {
		return err
	}

	if newBody.Error {
		return errors.New(newBody.Msg)
	}

	return nil
}