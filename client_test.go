package employmenthero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testClientID = "AXy9orp-CDaHhBZ9C78QHW2BKZpACgroqo85_NIOa9mIfJ9QnSVKzY-X_rivR_fTUUr6aLjcJsj6sDur"
var testSecret = "EBoIiUSkCKeSk49hHSgTem1qnjzzJgRQHDEHvGpzlLEf_nIoJd91xu8rPOBDCdR_UYNKVxJE-UgS2iCw"
var apiBaseSandbox = "https://google.com"

func TestNewClient(t *testing.T) {
	_, e := NewClient("", "", "")

	assert.Equal(t, e.Error(), "Client ID, Secret and APIBase are required to create a Client")

	c, _ := NewClient(testClientID, testSecret, apiBaseSandbox)

	assert.Equal(t, testClientID, c.ClientID)
	assert.Equal(t, testSecret, c.Secret)
	assert.Equal(t, apiBaseSandbox, c.APIBase)
}

