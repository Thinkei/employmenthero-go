package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type TaxDelcarationResponse struct {
	Data TaxDeclaration `json:"data"`
}

// Get returns a list of [TaxDeclaration] resources of one employee
//
// Example:
//
//	response, err := c.GetTaxDeclaration(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	taxDeclaration := response.Data
func (c *Client) GetTaxDeclaration(ctx context.Context, oid string, eid string, ep ListParams) (*TaxDelcarationResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/tax_declaration", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	response := &TaxDelcarationResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
