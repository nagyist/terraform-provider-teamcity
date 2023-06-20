package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id       *int64           `json:"id,omitempty"`
	Username string           `json:"username"`
	Password *string          `json:"password,omitempty"`
	Roles    *RoleAssignments `json:"roles,omitempty"`
}

type RoleAssignments struct {
	RoleAssignment []RoleAssignment `json:"role"`
}

type RoleAssignment struct {
	Id    string `json:"roleId"`
	Scope string `json:"scope"`
}

func (c *Client) NewUser(user User) (*User, error) {
	body, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/users", c.RestURL), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	result, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	actual := User{}
	err = json.Unmarshal(result, &actual)
	if err != nil {
		return nil, err
	}

	return &actual, nil
}

func (c *Client) GetUser(id string) (*User, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/users/id:%s", c.RestURL, id), nil)
	if err != nil {
		return nil, err
	}

	result, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	actual := User{}
	err = json.Unmarshal(result, &actual)
	if err != nil {
		return nil, err
	}

	return &actual, nil
}

func (c *Client) SetUser(user User) (*User, error) {
	rb, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/users/id:%d", c.RestURL, *user.Id), bytes.NewReader(rb))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	actual := User{}
	err = json.Unmarshal(body, &actual)
	if err != nil {
		return nil, err
	}

	return &actual, nil
}

func (c *Client) DeleteUser(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/users/id:%s", c.RestURL, id), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}