package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type SuperannuationDetailResponse struct {
	Data SuperannuationDetail `json:"data"`
}

// Get returns a list of [SuperannuationDetail] resources of one employee
//
// Example:
//
//	response, err := c.GetSuperannuationDetail(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
//	superannuationDetail := response.Data
func (c *Client) GetSuperannuationDetail(ctx context.Context, oid string, eid string, ep ListParams) (*SuperannuationDetailResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/superannuation_detail", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	response := &SuperannuationDetailResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
