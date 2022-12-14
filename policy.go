package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type PolicyListData struct {
	Items []Policy `json:"items"`
	ListResponse
}

type PolicyListResponse struct {
	Data PolicyListData `json:"data"`
}

func (c *Client) ListPolicies(ctx context.Context, oid string, tp ListParams) (*PolicyListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/policies", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &PolicyListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

