package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	config     ClientConfig
	httpClient *http.Client
}

type ClientConfig struct {
	ClientID     string
	ClientSecret string
	Subdomain    string
	Timeout      time.Duration
}

type AuthResponse struct {
	AccessToken  string    `json:"access_token,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	ExpiresIn    int       `json:"expires_in,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	TokenType    string    `json:"token_type,omitempty"`
	AccountID    int       `json:"account_id,omitempty"`
}

type method string

const (
	GET    method = http.MethodGet
	POST   method = http.MethodPost
	PUT    method = http.MethodPut
	DELETE method = http.MethodDelete
)

func NewClient(config ClientConfig) (*Client, error) {
	c := &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}

	// Attempt to authenticate
	_, err := c.getToken()

	return c, err
}

func (c *Client) getToken() (*AuthResponse, error) {
	authURL := fmt.Sprintf("https://%s.onelogin.com/auth/oauth2/v2/token", c.config.Subdomain)

	// Convert payload to JSON
	jsonData, err := json.Marshal(map[string]string{
		"grant_type": "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, authURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.config.ClientID, c.config.ClientSecret)
	req.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("authentication failed with status code %d", resp.StatusCode)
	}

	var authResponse AuthResponse
	err = json.NewDecoder(resp.Body).Decode(&authResponse)
	if err != nil {
		return nil, err
	}

	return &authResponse, err
}

func (c *Client) exec(method method, path string, body io.Reader, respModel interface{}) error {
	httpReq, err := http.NewRequest(string(method), fmt.Sprintf("https://%s.onelogin.com%s", c.config.Subdomain, path), body)
	if err != nil {
		return err
	}

	authResp, err := c.getToken()
	if err != nil {
		return err
	}

	httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authResp.AccessToken))
	httpReq.Header.Add("Content-Type", "application/json")

	resp, err := c.httpClient.Do(httpReq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("request failed with status code %d\n%s", resp.StatusCode, string(bodyBytes))
	}

	if respModel != nil {
		return json.NewDecoder(resp.Body).Decode(respModel)
	}

	return nil
}
