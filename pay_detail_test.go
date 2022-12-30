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

func TestListPayDetails(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "7d9e663f-a493-4523-aea4-4f299efeexxx", "effective_from": "2020-04-21", "classification": "Day Working - Full Time", "industrial_instrument": "Wholesale Award 2010", "pay_rate_template": "Permanent Storeworker", "anniversary_date": "2020-07-25T00:00:00+00:00", "salary": 50, "salary_type": "Hour", "pay_unit": "Hourly", "pay_category": "Permanent Ordinary Hours", "leave_allowance_template": "Permanent Leave Allowance", "change_reason": "Some reason...", "comments": "Some comments..." }],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListPayDetails(context.TODO(), "organisation_uid", "-", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []PayDetail{{Id: "7d9e663f-a493-4523-aea4-4f299efeexxx", EffectiveFrom: "2020-04-21", Classification: "Day Working - Full Time", IndustrialInstrument: "Wholesale Award 2010", PayRateTemplate: "Permanent Storeworker", AnniversaryDate: "2020-07-25T00:00:00+00:00", Salary: 50, SalaryType: "Hour", PayUnit: "Hourly", PayCategory: "Permanent Ordinary Hours", LeaveAllowanceTemplate: "Permanent Leave Allowance", ChangeReason: "Some reason...", Comments: "Some comments..."}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}

