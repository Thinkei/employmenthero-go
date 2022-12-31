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

func TestListCerification(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "6538a2bb-2b34-4437-9a00-92af3dab5b59", "name": "Full-disk encryption", "type": "check" }],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListCertifications(context.TODO(), "organisation_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Certification{{Id: "6538a2bb-2b34-4437-9a00-92af3dab5b59", Name: "Full-disk encryption", Type: "check"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
