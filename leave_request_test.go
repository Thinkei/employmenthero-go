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

func TestListLeaveRequest(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx","start_date":"2020-12-22T00:00:00+00:00","end_date":"2021-01-08T00:00:00+00:00","total_hours":83.6,"comment":"","status":"Declined","leave_balance_amount":0.0,"leave_category_name":"Annual Leave","reason":"test","employee_id":"xxxx-yyyyy"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListLeaveRequests(context.TODO(), "organisation_uid", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []LeaveRequest{{Id: "51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx", StartDate: "2020-12-22T00:00:00+00:00", EndDate: "2021-01-08T00:00:00+00:00", TotalHours: 83.6, Status: "Declined", LeaveCategoryName: "Annual Leave", Reason: "test", EmployeeId: "xxxx-yyyyy"}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

func TestGetLeaveRequest(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"id":"51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx","start_date":"2020-12-22T00:00:00+00:00","end_date":"2021-01-08T00:00:00+00:00","total_hours":83.6,"comment":"","status":"Declined","leave_balance_amount":0.0,"leave_category_name":"Annual Leave","reason":"test","employee_id":"xxxx-yyyyy"}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.GetLeaveRequest(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-xxxx")

	expectedResult := LeaveRequest{Id: "51c4b9c6-1ca5-4d72-8f75-6bb3a6xxxx", StartDate: "2020-12-22T00:00:00+00:00", EndDate: "2021-01-08T00:00:00+00:00", TotalHours: 83.6, Status: "Declined", LeaveCategoryName: "Annual Leave", Reason: "test", EmployeeId: "xxxx-yyyyy"}

	assert.Equal(t, response.Data, expectedResult)
	assert.Nil(t, err)
}
