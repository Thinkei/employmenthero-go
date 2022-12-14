package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type EmployeeCertificationList struct {
	Items []EmployeeCertification `json:"items"`
	ListResponse
}

type EmployeeCertificationListResponse struct {
	Data EmployeeCertificationList `json:"data"`
}

func (c *Client) ListEmployeeCertifications(ctx context.Context, oid string, eid string, hp ListParams) (*EmployeeCertificationListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/certifications", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", hp.PageIndex)
	q.Add("item_per_page", hp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &EmployeeCertificationListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

