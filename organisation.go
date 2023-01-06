package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type OrganisationData struct {
	Items []Organisation `json:"items"`
	ListResponse
}

type OrganisationListResponse struct {
	Data OrganisationData `json:"data"`
}

// Get returns a list of [Organisation] resources that the current user can access into
//
// Example:
//
//	response, err := c.ListOrganisations(context.TODO(), ListParams{})
//	organisations := response.Data.Items
func (c *Client) ListOrganisations(ctx context.Context, op ListParams) (*OrganisationListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations", c.APIBase), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", op.PageIndex)
	q.Add("item_per_page", op.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &OrganisationListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

type OrganisationResponse struct {
	Data OrganisationDetail `json:"data"`
}

// Get returns the details of one [Organisation] resource that the current user can access into
//
// Example:
//
//	response, err := c.GetOrganisation(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx")
//	organisation := response.Data
func (c *Client) GetOrganisation(ctx context.Context, oid string) (*OrganisationResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	response := &OrganisationResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
