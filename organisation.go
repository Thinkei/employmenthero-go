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

type OrganisationResponse struct {
	Data OrganisationData `json:"data"`
}

func (c *Client) ListOrganisations(ctx context.Context, op OrganisationListParams) (*OrganisationResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations", c.APIBase), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", op.PageIndex)
	q.Add("item_per_page", op.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &OrganisationResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
