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

func TestListOrganisation(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id":"3cfd1633-4920-xxxy-be7e-98i13159x74","name":"Employment Hero","phone":"+612803848123","country":"AU","logo_url":"https://logo.com"}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListOrganisations(context.TODO(), ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)

	expectedResult := []Organisation{{Id: "3cfd1633-4920-xxxy-be7e-98i13159x74", Name: "Employment Hero", Phone: "+612803848123", Country: "AU", LogoURL: "https://logo.com"}}

	assert.Equal(t, response.Data.Items, expectedResult)
	assert.Nil(t, err)
}

func TestShowOrganisation(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"id":"90a34ef1-50e4-4930-a9d6-xxxx","name":"Luke Luke","phone":"+61","country":"AU","logo_url":"http://logo.com","primary_address":"Opera House","end_of_week":"Sunday","typical_work_day":"7.6","payroll_admin_emails":[],"subscription_plan":"Demo","superfund_name":"REST","employees_count":10,"active_employees_count":10,"pending_employees_count":0,"business_account_id":null,"time_zone":"Australia/Sydney","created_at":"2019-09-12T23:59:30+01:00"}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.GetOrganisation(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx")

	expectedResult := OrganisationDetail{
		Id:                    "90a34ef1-50e4-4930-a9d6-xxxx",
		Name:                  "Luke Luke",
		Phone:                 "+61",
		TimeZone:              "Australia/Sydney",
		PrimaryAddress:        "Opera House",
		LogoURL:               "http://logo.com",
		Country:               "AU",
		EndOfWeek:             "Sunday",
		TypicalWorkDay:        "7.6",
		PayrollAdminEmails:    []string{},
		SubscriptionPlan:      "Demo",
		SuperfundName:         "REST",
		EmployeesCount:        10,
		ActiveEmployeesCount:  10,
		PendingEmployeesCount: 0,
		CreatedAt:             "2019-09-12T23:59:30+01:00",
	}

	assert.Equal(t, response.Data, expectedResult)
	assert.Nil(t, err)
}
