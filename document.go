package go_notify

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (n *Notify) CreateDocument(name string) (Document, error) {
	body, err := json.Marshal(map[string]string{
		"name": name,
	})

	if err != nil {
		return Document{}, err
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", fmt.Sprintf(`%s/document`, base_url), bytes.NewBuffer(body))

	if err != nil {
		return Document{}, err
	}

	req.Header.Set("x-api-key", n.Secret)

	resp, err := client.Do(req)

	if err != nil {
		return Document{}, err
	}

	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return Document{}, err
	}

	var new_body Response

	if err := json.Unmarshal(body, &new_body); err != nil {
		return Document{}, err
	}

	if new_body.Error {
		return Document{}, errors.New(new_body.Msg)
	}

	return new_body.Document, err
}

func (n *Notify) GetDocument(name string) (interface{}, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", fmt.Sprintf(`%s/document/%s`, base_url, name), nil)

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
