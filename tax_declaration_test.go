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

func TestGetTaxDeclaration(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"first_name":"www","last_name":"www","tax_file_number":"000000000","tax_au_resident":false,"tax_foreign_resident":true,"working_holiday_maker":false,"tax_free":false,"tax_help_debt":false,"tax_financial_supplement_debt":false}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.GetTaxDeclaration(context.TODO(), "organisation_uid", "xxx-yy", ListParams{})

	expectedResult := TaxDeclaration{FirstName: "www", LastName: "www", TaxFileNumber: "000000000", TaxForeignResident: true, WorkingHolidayMaker: false, TaxFree: false, TaxHelpDebt: false, TaxFinancialSupplementDebt: false}

	assert.Equal(t, response.Data, expectedResult)

	assert.Nil(t, err)
}

