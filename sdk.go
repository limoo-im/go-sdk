package sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/limoo-im/go-sdk/types"
)

// Main Client you need to create one as a pointer
type LimooClient struct {
	Username   string
	Password   string
	BaseURL    string
	httpClient *http.Client
	lastToken  *string
}

// login to limoo and return the client
// TODO: use refresh token
func (c *LimooClient) New(limooBaseURL, username, password string, insecureSkipVerify bool) error {
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
	c.httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecureSkipVerify,
		},
	}
	token, err := c.login()
	c.lastToken = token
	return err
}

// login to limoo and return the new token
// TODO: use refresh token
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

// send a message to a conversation
// TODO: return a readable response and also login if necessary
func (c *LimooClient) SendMessage(opts types.SendMessageOptions) error {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	var resError error
	reqUri := fmt.Sprintf("/Limonad/api/v1/workspace/items/%v/conversation/items/%v/message/items", opts.WorkspaceID, opts.ConversationID)
	res, err := c.httpClient.Post(c.BaseURL+reqUri, "application/json", bytes.NewReader(body))
	if res.StatusCode == http.StatusUnauthorized {
		resError = errors.New("bad username or password, check your credentials")
		return resError
	} else if res.StatusCode != http.StatusOK {
		resError = errors.New("unknown status code from server")
		return resError
	}
	return nil
}
