package employmenthero

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/Thinkei/employmenthero-go/mocks"
	"github.com/stretchr/testify/assert"
)

var testClientID = "AXy9orp-CDaHhBZ9C78QHW2BKZpACgroqo85"
var testSecret = "UgS2iCw"
var refreshToken = "w+12axYWx2UgS2iCw"
var oauthBase = "https://oauth.eh.com"
var apiBase = "https://api.eh.com"

var c *Client

func init() {
	c, _ = NewClient(testClientID, testSecret, refreshToken, oauthBase, apiBase)
	c.Client = &mocks.MockHttpClient{}
}

func TestNewClient(t *testing.T) {
	_, e := NewClient("", "", "", "", "")

	assert.Equal(t, e.Error(), "Client ID, Secret and APIBase are required to create a Client")

	assert.Equal(t, testClientID, c.ClientID)
	assert.Equal(t, testSecret, c.Secret)
	assert.Equal(t, refreshToken, c.Token.RefreshToken)
	assert.Equal(t, oauthBase, c.OAuthBase)
	assert.Equal(t, apiBase, c.APIBase)
}

func TestGetAccessTokenInvalidClientResponse(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"message":"Invalid token"}`)))
	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 400,
			Body: r,
		}, nil
	}
	response, err := c.GetAccessToken(context.TODO())
	assert.Empty(t, response.Token)
	assert.NotNil(t, err)
	assert.EqualValues(t, "{\"message\":\"Invalid token\"}", err.Error())
}

func TestGetAccessTokenSuccess(t *testing.T) {
	json := `{"access_token":"YYUUzz","refresh_token":"xxYYzz","token_type":"bearer","expires_in":900}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(json)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body: r,
		}, nil
	}
	response, err := c.GetAccessToken(context.TODO())
	assert.Nil(t, err)
	assert.Equal(t, response.RefreshToken, "xxYYzz")
	assert.Equal(t, response.Token, "YYUUzz")
	assert.Equal(t, response.Type, "bearer")
	assert.Equal(t, response.ExpiresIn, 900)
}
