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

func TestListBankAccounts(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"items":[{"id": "883280d7-470c-42e0-83b2", "account_name":"peter","account_number":"000000000","bsb":"012345","amount":"0.0","primary_account":true}],"item_per_page":20,"page_index":1,"total_pages":1,"total_items":1}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.ListBankAccounts(context.TODO(), "organisation_uid", "9e080b45-244f-4fbd", ListParams{})

	assert.Equal(t, response.Data.ItemPerPage, 20)
	assert.Equal(t, response.Data.PageIndex, 1)
	assert.Equal(t, response.Data.TotalItems, 1)
	assert.Equal(t, response.Data.TotalPages, 1)
	expectedResult := []BankAccount{{Id: "883280d7-470c-42e0-83b2", AccountName: "peter", AccountNumber: "000000000", Bsb: "012345", Amount: "0.0", PrimaryAccount: true}}

	assert.Equal(t, response.Data.Items, expectedResult)

	assert.Nil(t, err)
}
