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

func TestListOrganisation(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"3cfd1633-4920-xxxy-be7e-98i13159x74","name":"Employment Hero","phone":"+612803848123","country":"AU","logo_url":"https://logo.com"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListOrganisations(context.TODO(), OrganisationListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)

	expectedResult := []Organisation{{ Id: "3cfd1633-4920-xxxy-be7e-98i13159x74", Name: "Employment Hero", Phone: "+612803848123", Country: "AU", LogoURL: "https://logo.com" }}

	assert.Equal(t, response.Data.Items, expectedResult)
	assert.Nil(t, err)
}
