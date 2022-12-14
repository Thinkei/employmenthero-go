package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type CertificationListData struct {
	Items []Certification `json:"items"`
	ListResponse
}

type CertificationListResponse struct {
	Data CertificationListData `json:"data"`
}

func (c *Client) ListCertifications(ctx context.Context, oid string, tp ListParams) (*CertificationListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/certifications", c.APIBase, oid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", tp.PageIndex)
	q.Add("item_per_page", tp.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &CertificationListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}

