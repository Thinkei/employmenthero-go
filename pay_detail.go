package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type PayDetailListData struct {
	Items []PayDetail `json:"items"`
	ListResponse
}

type PayDetailListResponse struct {
	Data PayDetailListData `json:"data"`
}

// Get returns a list of [PayDetail] resources of one employee
//
// Example:
//
//	response, err := c.ListPayDetails(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	payDetails := response.Data.Items
func (c *Client) ListPayDetails(ctx context.Context, oid string, eid string, tp ListParams) (*PayDetailListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/pay_details", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &PayDetailListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
