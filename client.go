package employmenthero

import "errors"

type Client struct {
	ClientID string
	Secret string
	APIBase string
}

// NewClient returns a new Client struct
func NewClient(clientID string, secret string, APIBase string) (*Client, error) {
	if clientID == "" || secret == "" || APIBase == "" {
		return nil, errors.New("Client ID, Secret and APIBase are required to create a Client")
	}

	return &Client{
		ClientID: clientID,
		Secret: secret,
		APIBase: APIBase,
	}, nil
}

