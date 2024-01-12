package onelogin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	urlpkg "net/url"
	"strconv"
	"time"
)

const (
	DefaultTimeout = 10 * time.Second
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

// Paging holds common paging parameters for APIs that support paging
// https://developers.onelogin.com/api-docs/2/getting-started/using-query-parameters#pagination
type Paging struct {
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
	Cursor string `json:"cursor,omitempty"`
}

func NewClient(config ClientConfig) (*Client, error) {
	if config.Timeout == 0 {
		config.Timeout = DefaultTimeout
	}

	c := &Client{
		config: config,
		httpClient: &http.Client{
			Timeout: config.Timeout,
		},
	}

	// Attempt to authenticate
	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	_, err := c.getToken(ctx)

	return c, err
}

func (c *Client) getToken(ctx context.Context) (*AuthResponse, error) {
	authURL := fmt.Sprintf("https://%s.onelogin.com/auth/oauth2/v2/token", c.config.Subdomain)

	// Convert payload to JSON
	jsonData, _ := json.Marshal(map[string]string{
		"grant_type": "client_credentials",
	})

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, authURL, bytes.NewBuffer(jsonData))
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
	return c.execRequest(
		&oneloginRequest{
			method:    method,
			path:      path,
			body:      body,
			respModel: respModel,
		},
	)

}

type oneloginRequest struct {
	method      method
	path        string
	body        io.Reader
	queryParams map[string]string
	respModel   interface{}
}

func (c *Client) execRequest(req *oneloginRequest) error {
	return c.execRequestContext(context.Background(), req)
}

func (c *Client) execRequestContext(ctx context.Context, req *oneloginRequest) error {
	// add configured timeout to context
	ctx, cancel := context.WithTimeout(ctx, c.config.Timeout)
	defer cancel()

	url := fmt.Sprintf("https://%s.onelogin.com%s", c.config.Subdomain, req.path)
	if req.queryParams != nil && len(req.queryParams) > 0 {
		queryParams := urlpkg.Values{}
		for key, value := range req.queryParams {
			queryParams.Add(key, value)
		}

		url += "?" + queryParams.Encode()
	}

	httpReq, err := http.NewRequestWithContext(ctx, string(req.method), url, req.body)
	if err != nil {
		return err
	}

	authResp, err := c.getToken(ctx)
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

	if req.respModel != nil {
		return json.NewDecoder(resp.Body).Decode(req.respModel)
	}

	return nil
}

func addPagingParams(queryParams map[string]string, paging *Paging) map[string]string {
	if paging.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(paging.Limit)
	}
	if paging.Page > 0 {
		queryParams["page"] = strconv.Itoa(paging.Page)
	}
	if paging.Cursor != "" {
		queryParams["cursor"] = paging.Cursor
	}

	return queryParams
}
