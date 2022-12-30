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

func TestListEmergencyContacts(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{ "id": 123, "contact_name": "Daniel Levi"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListEmergencyContacts(context.TODO(), "organisation_uid", "9e080b45-244f-4fbd", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []EmergencyContact{{Id: 123, ContactName: "Daniel Levi"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

