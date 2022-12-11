package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type TimesheetListData struct {
	Items []Timesheet `json:"items"`
	ListResponse
}

type TimesheetListResponse struct {
	Data TimesheetListData `json:"data"`
}

func (c *Client) ListTimesheetEntries(ctx context.Context, oid string, eid string, tp ListParams) (*TimesheetListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/timesheet_entries", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &TimesheetListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

