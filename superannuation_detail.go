package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type SuperannuationDetailResponse struct {
	Data SuperannuationDetail `json:"data"`
}

func (c *Client) GetSuperannuationDetail(ctx context.Context, oid string, eid string, ep ListParams) (*SuperannuationDetailResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/superannuation_detail", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	response := &SuperannuationDetailResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

