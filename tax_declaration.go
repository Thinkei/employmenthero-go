package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type TaxDelcarationResponse struct {
	Data TaxDeclaration `json:"data"`
}

func (c *Client) GetTaxDeclaration(ctx context.Context, oid string, eid string, ep ListParams) (*TaxDelcarationResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/tax_declaration", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	response := &TaxDelcarationResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

