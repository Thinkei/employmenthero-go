package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type PayslipList struct {
	Items []Payslip `json:"items"`
	ListResponse
}

type PayslipListReponse struct {
	Data PayslipList `json:"data"`
}

// Get returns a list of [Payslip] resources
//
// Example:
//
//	response, err := c.ListPayslips(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	payslips := response.Data.Items
func (c *Client) ListPayslips(ctx context.Context, oid string, eid string, hp ListParams) (*PayslipListReponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/payslips", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", hp.PageIndex)
	q.Add("item_per_page", hp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &PayslipListReponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
