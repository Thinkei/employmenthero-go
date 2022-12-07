package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type OrganisationListParams struct {
	ListParams
}

type OrganisationData struct {
	Items []Organisation `json:"items"`
	ListResponse
}

type OrganisationListResponse struct {
	Data OrganisationData `json:"data"`
}

func (c *Client) ListOrganisations(ctx context.Context, op OrganisationListParams) (*OrganisationListResponse, error) {
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

func (c *Client) GetOrganisation(ctx context.Context, oid string) (*OrganisationResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	response := &OrganisationResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
