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

func TestListTeams(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx", "name": "Developers", "status": "active"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListTeams(context.TODO(), "organisation_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Team{{Id: "51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx", Name: "Developers", Status: "active"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

func TestListEmployeesByTeam(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListEmployeesByTeam(context.TODO(), "organisation_uid", "team_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Employee{{Id: "51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
