package employmenthero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testClientID = "AXy9orp-CDaHhBZ9C78QHW2BKZpACgroqo85"
var testSecret = "UgS2iCw"
var refreshToken = "w+12axYWx2UgS2iCw"
var apiBaseSandbox = "https://google.com"

func TestNewClient(t *testing.T) {
	_, e := NewClient("", "", "", "")

	assert.Equal(t, e.Error(), "Client ID, Secret and APIBase are required to create a Client")

	c, _ := NewClient(testClientID, testSecret, refreshToken, apiBaseSandbox)

	assert.Equal(t, testClientID, c.ClientID)
	assert.Equal(t, testSecret, c.Secret)
	assert.Equal(t, refreshToken, c.Token.RefreshToken)
	assert.Equal(t, apiBaseSandbox, c.APIBase)
}
