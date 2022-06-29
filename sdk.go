package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/limoo-im/go-sdk/types"
)

type LimooClient struct {
	Username   string
	Password   string
	BaseURL    string
	httpClient *http.Client
}

// login to limoo and return the client
func (c *LimooClient) New(limooBaseURL, username, password string) error {
	if c != nil {
		c.Username = username
		c.Password = password
		c.BaseURL = limooBaseURL
		c.httpClient = &http.Client{}
	} else {
		c = &LimooClient{
			Username:   username,
			Password:   password,
			httpClient: &http.Client{},
			BaseURL:    limooBaseURL,
		}
	}
	_, err := c.login()
	return err
}

func (c *LimooClient) login() (*string, error) {
	body := fmt.Sprintf("j_username=%v&j_password=%v", c.Username, c.Password)
	res, err := c.httpClient.Post(c.BaseURL+"/Limonad/j_spring_security_check", "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	var resError error
	if res.StatusCode == http.StatusUnauthorized {
		resError = errors.New("bad username or password, check your credentials")
		return nil, resError
	} else if res.StatusCode != http.StatusOK {
		resError = errors.New("unknown status code from server")
		return nil, resError
	}
	respbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var loginResp types.LoginResult
	err = json.Unmarshal(respbody, &loginResp)
	if err != nil {
		return nil, err
	}
	token := res.Header.Get("Token")
	return &token, nil
}
