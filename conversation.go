package robin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (r *Robin) CreateConversation(senderName, senderId, receiverId, receiverName string) (interface{}, error) {
	body, err := json.Marshal(map[string]string{
		"sender_name":   senderName,
		"sender_id":     senderId,
		"receiver_id":   receiverId,
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

	var newBody Response

	if err := json.Unmarshal(body, &newBody); err != nil {
		return nil, err
	}

	if newBody.Error {
		return nil, errors.New(newBody.Msg)
	}

	return newBody.Data, nil
}

func (r *Robin) GetConversationMessages(id string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(`%s/conversation/messages/%s`, baseUrl, id), nil)

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

	var new_body Response

	if err := json.Unmarshal(body, &new_body); err != nil {
		return nil, err
	}

	if new_body.Error {
		return nil, errors.New(new_body.Msg)
	}

	return new_body.Data, nil
}
