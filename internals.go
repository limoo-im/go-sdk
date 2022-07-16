package sdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/limoo-im/go-sdk/types"
	log "github.com/sirupsen/logrus"
)

// login to limoo and return the new token
func (c *LimooClient) login() error {
	body := fmt.Sprintf("j_username=%v&j_password=%v", c.Username, c.Password)
	res, err := c.httpClient.Post(c.BaseURL+"/Limonad/j_spring_security_check", "application/x-www-form-urlencoded", strings.NewReader(body))
	if err != nil {
		return err
	}
	var resError error
	if res.StatusCode == http.StatusUnauthorized {
		resError = errors.New("bad username or password, check your credentials")
		return resError
	} else if res.StatusCode != http.StatusOK {
		resError = errors.New("unknown status code from server")
		return resError
	}
	respbody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	var loginResp types.LoginResult
	err = json.Unmarshal(respbody, &loginResp)
	if err != nil {
		return err
	}
	token := res.Header.Get("Token")
	c.lastToken = &token
	// TODO: Log response body
	log.WithField("event", "login").Debugf("Headers: %v", res.Header)
	log.WithField("event", "login").Debugf("Login Data: %v", loginResp)
	return nil
}

// a function that is used to send all of the requests (except for the login) and login again if gets Unauthorized status
func (c *LimooClient) do(uri string, method, contentType string, body io.Reader) (*http.Response, error) {
	url, err := url.Parse(c.BaseURL + uri)
	if err != nil {
		return nil, err
	}
	request := &http.Request{
		Method: method,
		URL:    url,
		Body:   io.NopCloser(body),
		Header: http.Header{
			"Content-Type":  []string{contentType},
			"Authorization": []string{fmt.Sprintf("Bearer %s", *c.lastToken)},
		},
	}
	log.WithField("event", "request to server").Debugf("request info: %+v", *request)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode == http.StatusUnauthorized {
		err = c.login()
		if err != nil {
			return nil, err
		}
		response, err = c.httpClient.Do(request)
		if err != nil {
			return nil, err
		}
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unknown status %v code from server", response.StatusCode)
	}
	log.WithField("event", "response from server").Debugf("Headers: %v", response.Header)
	return response, nil
}
