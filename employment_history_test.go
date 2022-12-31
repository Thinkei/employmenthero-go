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

func TestListEmploymentHistories(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{ "id": "0f551b48-c958-4359-87c2", "title": "Grad developer", "start_date":"2017-03-01T00:00:00+00:00", "end_date": null, "employment_type": "Part-time" }],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListEmploymentHistories(context.TODO(), "organisation_uid", "9e080b45-244f-4fbd", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []EmploymentHistory{{Id: "0f551b48-c958-4359-87c2", Title: "Grad developer", StartDate: "2017-03-01T00:00:00+00:00", EndDate: "", EmploymentType: "Part-time"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
