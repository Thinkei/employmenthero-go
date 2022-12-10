package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)


type LeaveRequestListData struct {
	Items []LeaveRequest `json:"items"`
	ListResponse
}

type LeaveRequestListResponse struct {
	Data LeaveRequestListData `json:"data"`
}

func (c *Client) ListLeaveRequests(ctx context.Context, oid string, ep ListParams) (*LeaveRequestListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/leave_requests", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", ep.PageIndex)
	q.Add("item_per_page", ep.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &LeaveRequestListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

type LeaveRequestResponse struct {
	Data LeaveRequest `json:"data"`
}

func (c *Client) GetLeaveRequest(ctx context.Context, oid string, lid string) (*LeaveRequestResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/leave_requests/%s", c.APIBase, oid, lid), nil)

	if err != nil {
		return nil, err
	}

	response := &LeaveRequestResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

