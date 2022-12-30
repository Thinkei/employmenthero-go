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
	"time"
)

// NewClient returns a new Client struct
func NewClient(clientID, secret, redirectUri, OAuthBase, APIBase string) (*Client, error) {
	if clientID == "" || secret == "" || APIBase == "" || OAuthBase == "" || redirectUri == "" {
		return nil, errors.New("Client ID, Secret, RedirectURI, OAuthBase and APIBase are required to create a Client")
	}

	return &Client{
		Client:   &http.Client{},
		ClientID: clientID,
		RedirectURI: redirectUri,
		Secret:   secret,
		APIBase:  APIBase,
		OAuthBase:  OAuthBase,
	}, nil
}

// SetLog will set/change the output destination.
func (c *Client) SetLog(log io.Writer) {
	c.Log = log
}

func (c *Client) GetOAuth2Access(ctx context.Context, code string) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.Secret)
	data.Set("code", code)
	data.Set("redirect_uri", "https://google.com/callbacck")
	data.Set("grant_type", "authorization_code")

	buf := bytes.NewBuffer([]byte(data.Encode()))
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.OAuthBase, "/oauth2/token"), buf)
	if err != nil {
		return &TokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := &TokenResponse{}

	err = c.Send(req, response)

	if response.Token != "" {
		c.Token = response
		c.tokenExpiresAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
	}

	return response, err
}

func (c *Client) SetRefreshToken(refreshToken string) {
	c.Token = &TokenResponse{RefreshToken: refreshToken}
}

func (c *Client) GetAccessToken(ctx context.Context) (*TokenResponse, error) {
	data := url.Values{}
	data.Set("client_id", c.ClientID)
	data.Set("client_secret", c.Secret)
	data.Set("refresh_token", c.Token.RefreshToken)
	data.Set("grant_type", "refresh_token")

	buf := bytes.NewBuffer([]byte(data.Encode()))
	req, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s%s", c.OAuthBase, "/oauth2/token"), buf)
	if err != nil {
		return &TokenResponse{}, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response := &TokenResponse{}

	err = c.Send(req, response)

	if response.Token != "" {
		c.Token = response
		c.tokenExpiresAt = time.Now().Add(time.Duration(response.ExpiresIn) * time.Second)
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
	c.log(req, resp)

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

func (c *Client) SendWithAuth(req *http.Request, v interface{}) error {
	if c.Token != nil {
		if !c.tokenExpiresAt.IsZero() && time.Until(c.tokenExpiresAt) < RequestNewTokenBeforeExpiresIn {
			if _, err := c.GetAccessToken(req.Context()); err != nil {
				return err
			}
		}

		req.Header.Set("Authorization", "Bearer "+c.Token.Token)
	}

	return c.Send(req, v)
}

func (c *Client) NewRequest(ctx context.Context, method, url string, payload interface{}) (*http.Request, error) {
	var buf io.Reader
	if payload != nil {
		b, err := json.Marshal(&payload)
		if err != nil {
			return nil, err
		}

		buf = bytes.NewBuffer(b)
	}
	return http.NewRequestWithContext(ctx, method, url, buf)
}

func (c *Client) log(r *http.Request, resp *http.Response) {
	if c.Log == nil { return }

	var (
		reqDump string
		respDump []byte
	)

	if r != nil {
		reqDump = fmt.Sprintf("%s %s", r.Method, r.URL.String())
	}

	c.Log.Write([]byte(fmt.Sprintf("Request: %s\nResponse: %s\n", reqDump, string(respDump))))
}

