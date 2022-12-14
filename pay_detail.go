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

