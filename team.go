package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type TeamListData struct {
	Items []Team `json:"items"`
	ListResponse
}

type TeamListDataResponse struct {
	Data TeamListData `json:"data"`
}

// Get returns a list of [Team] resources
//
// Example:
//
//	response, err := c.ListTeams(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
//	teams := response.Data.Items
func (c *Client) ListTeams(ctx context.Context, oid string, tp ListParams) (*TeamListDataResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/teams", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &TeamListDataResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

// Get returns a list of [Employee] resources which belong to one [Team]
//
// Example:
//
//	response, err := c.ListEmployeesByTeam(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXXX-YYYY-ZZZZ", ListParams{})
//	employees := response.Data.Items
func (c *Client) ListEmployeesByTeam(ctx context.Context, oid string, tid string, tp ListParams) (*EmployeeListReponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/teams/%s/employees", c.APIBase, oid, tid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &EmployeeListReponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
