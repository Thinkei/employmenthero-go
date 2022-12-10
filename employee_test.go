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

func init() {
	c, _ = NewClient(testClientID, testSecret, refreshToken, oauthBase, apiBase)
	c.Client = &mocks.MockHttpClient{}
}

func TestListEmployees(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"3cfd1633-4920-xxxy-be7e-98i13159x74", "account_email": "dev@employmenthero.com", "title": "dev", "role": "employee", "first_name":"Hoa", "last_name": "Nguyen", "Address": "Sydney, NSW"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListEmployees(context.TODO(), "organisation_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Employee{{Id: "3cfd1633-4920-xxxy-be7e-98i13159x74", AccountEmail: "dev@employmenthero.com", Title: "dev", Role: "employee", FirstName: "Hoa", LastName: "Nguyen", Address: "Sydney, NSW"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
