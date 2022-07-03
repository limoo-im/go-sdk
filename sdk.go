package sdk

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/limoo-im/go-sdk/types"
	log "github.com/sirupsen/logrus"
)

// Main Client you need to create one as a pointer
type LimooClient struct {
	Username   string
	Password   string
	BaseURL    string
	httpClient *http.Client
	lastToken  *string
}

// print debug messages or not? (will print sensitive information)
func SetDebug(d bool) {
	if d {
		log.SetLevel(log.DebugLevel)
	}
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
	return c.login()
}

// send a message to a conversation
// TODO: return a readable response
func (c *LimooClient) SendMessage(opts types.SendMessageOptions) error {
	body, err := json.Marshal(opts)
	if err != nil {
		return err
	}
	reqUri := fmt.Sprintf("/Limonad/api/v1/workspace/items/%v/conversation/items/%v/message/items", opts.WorkspaceID, opts.ConversationID)
	res, err := c.do(reqUri, http.MethodPost, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	// TODO: Log response body
	log.WithField("event", "send message").Debugf("Headers: %v", res.Header)
	return nil
}
