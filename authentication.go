package forlabs

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/url"
	"strconv"
)

type AuthenticateObject struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Authenticate method performs a request to https://bki.forlabs.ru/app/login and if successful, returns Client instance
// https://bki.forlabs.ru/app/login JSON username,password
func Authenticate(username, password string) (c *Client, err error) {
	client := &Client{
		HTTPClient: resty.New(),
	}

	auth := AuthenticateObject{username, password}

	// Obtain XSRF token
	resp, err := client.HTTPClient.R().Get(Endpoint+"/app/login")
	if err != nil {
		return nil, err
	}
	cookies := resp.Cookies()
	for _, v := range cookies {
		if v.Name == "forlabs_session" {
			client.CurrentToken, err = url.QueryUnescape(v.Value)
			if err != nil {
				return
			}
		}
		if v.Name == "XSRF-TOKEN" {
			client.XSRFToken, err = url.QueryUnescape(v.Value)
			if err != nil {
				return
			}
		}
	}
	if client.CurrentToken == "" || client.XSRFToken == "" {
		return nil, errors.New("failed to obtain both tokens")
	}
	resp, err = client.Post(Endpoint+"/app/login", auth)
	if err != nil {
		return nil, err
	}
	code := resp.StatusCode()
	if code == 422 {
		return nil, errors.New("invalid password")
	} else if code != 200 {
		return nil, errors.New("invalid status code " + strconv.Itoa(code))
	}
	client.Username = username
	client.Password = password
	return client, nil
}