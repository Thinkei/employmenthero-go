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

// Get returns a list of [EmployeeCertification] resources of one employee
//
// Example:
//
//	response, err := c.ListEmployeeCertifications(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	employeeCertifications := response.Data.Items
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
