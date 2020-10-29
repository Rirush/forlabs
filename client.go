package forlabs

import (
	"github.com/go-resty/resty/v2"
	"net/http"
	"net/url"
	"sync"
)

const Endpoint = "https://bki.forlabs.ru"

type Client struct {
	// Credentials are stored for re-authentication in case token expires
	Username string
	Password string

	// Tokens are changing after each request, so they are stored with a mutex
	TokenMutex sync.Mutex
	CurrentToken string
	XSRFToken string

	HTTPClient *resty.Client
}

//func (c *Client) Get(path string) (resp *resty.Response, err error) {
//
//}

func (c *Client) Post(path string, data interface{}) (resp *resty.Response, err error) {
	c.TokenMutex.Lock()
	defer c.TokenMutex.Unlock()
	resp, err = c.HTTPClient.R().
		SetCookie(&http.Cookie{Name: "forlabs_session", Value: c.CurrentToken}).
		SetCookie(&http.Cookie{Name: "XSRF-TOKEN", Value: c.XSRFToken}).
		SetContentLength(true).
		SetHeader("x-xsrf-token", c.XSRFToken).
		SetHeader("Accept", "application/json").
		SetHeader("Content-Type", "application/json").
		SetBody(data).
		Post(path)
	if err != nil {
		return
	}
	// TODO: handle error 419 (CSRF expiration)

	for _, v := range resp.Cookies() {
		if v.Name == "forlabs_session" {
			c.CurrentToken, err = url.QueryUnescape(v.Value)
			if err != nil {
				return
			}
		}
		if v.Name == "XSRF-TOKEN" {
			c.XSRFToken, err = url.QueryUnescape(v.Value)
			if err != nil {
				return
			}
		}
	}
	return
}