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

func TestListTimesheetEntries(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"0f551b48-c958-4359-87c2","employee_id":"9e080b45-244f-4fbd-88d1","date":"2020-12-03T00:00:00+00:00","start_time":null,"end_time":null,"status":"approved","units":5.0,"reason":null,"comment":"h","time":18000,"cost_centre":{"id":"e9a4df80-5d05-444b-ab09","name":"Hoa's Bakery"}}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListTimesheetEntries(context.TODO(), "organisation_uid", "-", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Timesheet{{Id: "0f551b48-c958-4359-87c2", EmployeeId: "9e080b45-244f-4fbd-88d1", Date: "2020-12-03T00:00:00+00:00", StartTime: "", EndTime: "", Comment: "h", Reason: "", Time: 18000, CostCentre: BasicData{ Id: "e9a4df80-5d05-444b-ab09", Name: "Hoa's Bakery" }, Status: "approved", Units: 5.0 }}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

