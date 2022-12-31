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

func TestGetSuperannuationDetail(t *testing.T) {
	r := ioutil.NopCloser(bytes.NewReader([]byte(`{"data":{"fund_name": "SUPER Super", "member_number": "184752", "product_code": "29400111", "employer_nominated_fund": true, "fund_abn": "19905421111", "electronic_service_address": "CLICKSUPER", "fund_email": "some_email@email.com", "account_name": "Daniel", "account_bsb": "HST011xxx", "account_number": "25767731xxx"}}`)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	response, err := c.GetSuperannuationDetail(context.TODO(), "organisation_uid", "xxx-yy", ListParams{})

	expectedResult := SuperannuationDetail{
		FundName:                 "SUPER Super",
		MemberNumber:             "184752",
		ProductCode:              "29400111",
		EmployerNominatedFund:    true,
		FundAbn:                  "19905421111",
		ElectronicServiceAddress: "CLICKSUPER",
		FundEmail:                "some_email@email.com",
		AccountName:              "Daniel",
		AccountBSB:               "HST011xxx",
		AccountNumber:            "25767731xxx",
	}

	assert.Equal(t, response.Data, expectedResult)

	assert.Nil(t, err)
}
