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

func TestListEmployeeCertification(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "3128a2bb-2b34-4437-9a00", "name": "Full-disk encryption", "certification_id": "6538a2bb-2b34-4437-9a00", "type": "Check", "expiry_date": "2021-01-04T00:00:00+00:00", "completion_date": "2020-01-04T00:00:00+00:00", "driver_problem": false, "driver_details": "", "status": "Outstanding", "reason": "" }],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListEmployeeCertifications(context.TODO(), "organisation_uid", "9e080b45-244f-4fbd", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []EmployeeCertification{{
		Id: "3128a2bb-2b34-4437-9a00",
		Name: "Full-disk encryption",
		CertificationId: "6538a2bb-2b34-4437-9a00",
		Type: "Check",
		ExpiryDate: "2021-01-04T00:00:00+00:00",
		CompletionDate: "2020-01-04T00:00:00+00:00",
		DriverProblem: false,
		Status: "Outstanding",
		Reason: "",
	}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

