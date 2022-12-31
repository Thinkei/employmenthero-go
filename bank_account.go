package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type BankAccountList struct {
	Items []BankAccount `json:"items"`
	ListResponse
}

type BankAccountListReponse struct {
	Data BankAccountList `json:"data"`
}

// Get returns a list of [BankAccount] resources of one employee
//
// Example:
//
//	response, err := c.ListBankAccounts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	bankAccounts := response.Data.Items
func (c *Client) ListBankAccounts(ctx context.Context, oid string, eid string, ep ListParams) (*BankAccountListReponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/bank_accounts", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", ep.PageIndex)
	q.Add("item_per_page", ep.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &BankAccountListReponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
