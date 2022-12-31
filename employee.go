package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type EmployeeListData struct {
	Items []Employee `json:"items"`
	ListResponse
}

type EmployeeListReponse struct {
	Data EmployeeListData `json:"data"`
}

// Get returns a list of [Employee] resources
//
// Example:
//
//	response, err := c.ListEmployees(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
//	employees := response.Data.Items
func (c *Client) ListEmployees(ctx context.Context, oid string, ep ListParams) (*EmployeeListReponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", ep.PageIndex)
	q.Add("item_per_page", ep.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &EmployeeListReponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

type EmployeeResponse struct {
	Data Employee `json:"data"`
}

// Get returns the details of an [Employee] resource
//
// Example:
//
//	response, err := c.GetEmployee(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-yyyy")
//	employee := response.Data
func (c *Client) GetEmployee(ctx context.Context, oid string, eid string) (*EmployeeResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	response := &EmployeeResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
