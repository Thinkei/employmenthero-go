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

func TestListPayslips(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "e27387ba-2105-4d12-be95", "first_name": "Daniel", "last_name": "Nguyen", "total_deduction": null, "net_pay": 2520.15, "super": 301.15, "wages": 3346.15, "reimbursements": null, "tax": 826.0, "name": "Alex Kopczynski", "address_line1": "4 Sava", "address_line2": null, "suburb": "123 sydney", "postcode": "0880", "post_tax_deduction": 0.0, "pre_tax_deduction": 0.0, "base_rate": 87000.0, "hourly_rate": 0.0, "base_rate_unit": "Annually", "employment_type": "Annually", "payment_date": "2012-07-11T00:00:00+10:00"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListPayslips(context.TODO(), "organisation_uid", "9e080b45-244f-4fbd", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []Payslip{{
		Id:               "e27387ba-2105-4d12-be95",
		FirstName:        "Daniel",
		LastName:         "Nguyen",
		TotalDeduction:   0.0,
		NetPay:           2520.15,
		Super:            301.15,
		Wages:            3346.15,
		Reimbursements:   0.0,
		Tax:              826.0,
		Name:             "Alex Kopczynski",
		AddressLine1:     "4 Sava",
		AddressLine2:     "",
		Suburb:           "123 sydney",
		Postcode:         "0880",
		PostTaxDeduction: 0.0,
		PreTaxDeduction:  0.0,
		BaseRate:         87000.0,
		HourlyRate:       0.0,
		BaseRateUnit:     "Annually",
		EmploymentType:   "Annually",
		PaymentDate:      "2012-07-11T00:00:00+10:00",
	}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
