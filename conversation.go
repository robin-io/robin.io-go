package go_notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (n *Notify) CreateConversation(sender_name, sender_id, reciever_id, reciever_name string) (interface{}, error) {
	body, err := json.Marshal(map[string]string{
		"sender_name":   sender_name,
		"sender_id":     sender_id,
		"reciever_id":   reciever_id,
		"reciever_name": reciever_name,
	})

	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(`%s/conversation`, base_url), bytes.NewBuffer(body))

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", n.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
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

func (n *Notify) GetConversationMessages(id string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(`%s/conversation/messages/%s`, base_url, id), nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("x-api-key", n.Secret)

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
