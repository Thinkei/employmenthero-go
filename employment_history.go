package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type EmploymentHistoryList struct {
	Items []EmploymentHistory `json:"items"`
	ListResponse
}

type EmploymentHistoryListResponse struct {
	Data EmploymentHistoryList `json:"data"`
}

func (c *Client) ListEmploymentHistories(ctx context.Context, oid string, eid string, hp ListParams) (*EmploymentHistoryListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/employment_histories", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", hp.PageIndex)
	q.Add("item_per_page", hp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &EmploymentHistoryListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

