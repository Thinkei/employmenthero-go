package employmenthero

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	ClientID    string
	Secret      string
	RedirectURI string
	APIBase     string
	Client      *http.Client
	Token       *TokenResponse
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	Token        string `json:"access_token"`
	Type         string `json:"token_type"`
	ExpiresIn    int64  `json:"expires_in"`
}

type ErrorResponse struct {
	Response *http.Response
}

// NewClient returns a new Client struct
func NewClient(clientID string, secret string, refreshToken string, APIBase string) (*Client, error) {
	if clientID == "" || secret == "" || APIBase == "" || refreshToken == "" {
		return nil, errors.New("Client ID, Secret and APIBase are required to create a Client")
	}

	return &Client{
		Client:   &http.Client{},
		ClientID: clientID,
		Secret:   secret,
		Token:    &TokenResponse{RefreshToken: refreshToken},
		APIBase:  APIBase,
	}, nil
}

func (c *Client) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.Secret)
	data.Set("refresh_token", c.Token.RefreshToken)
	data.Set("grant_type", "refresh_token")

	buf := bytes.NewBuffer([]byte(data.Encode()))
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.APIBase, "/oauth2/token"), buf)
	if err != nil {
		return &TokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := &TokenResponse{}

	err = c.Send(req, response)

	if response.Token != "" {
		c.Token = response
	}

	return response, err
}

func (c *Client) Send(req *http.Request, v interface{}) error {
	if v == nil {
		return nil
	}

	var (
		err  error
		resp *http.Response
	)

	req.Header.Set("Accept", "application/json")

	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err = c.Client.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) error {
		return Body.Close()
	}(resp.Body)

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		data, err := io.ReadAll(resp.Body)

		if err == nil && len(data) > 0 {
			return fmt.Errorf(string(data))
		}

	}

	if n, e := v.(io.Writer); e {
		_, err := io.Copy(n, resp.Body)
		return err
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
