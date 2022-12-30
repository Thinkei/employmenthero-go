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

func TestListPolicy(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "1d25a390-7b0e-0135-0209", "name": "Wonderful Company - Full-Disk Encryption Policy", "induction": true, "created_at": "2017-09-14T10:03:51+10:00"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListPolicies(context.TODO(), "organisation_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Policy{{Id: "1d25a390-7b0e-0135-0209", Name: "Wonderful Company - Full-Disk Encryption Policy", Induction: true, CreatedAt: "2017-09-14T10:03:51+10:00"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

