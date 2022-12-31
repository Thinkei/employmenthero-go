package employmenthero

import (
	"context"
	"fmt"
	"net/http"
)

type EmergencyContactList struct {
	Items []EmergencyContact `json:"items"`
	ListResponse
}

type EmergencyContactListResponse struct {
	Data EmergencyContactList `json:"data"`
}

// Get returns a list of [EmergencyContact] resources of one employee
//
// Example:
//
//	response, err := c.ListEmergencyContacts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "xxxx-yyy", ListParams{})
//	contacts := response.Data.Items
func (c *Client) ListEmergencyContacts(ctx context.Context, oid string, eid string, ep ListParams) (*EmergencyContactListResponse, error) {
	req, err := c.NewRequest(ctx, http.MethodGet, fmt.Sprintf("%s/api/v1/organisations/%s/employees/%s/emergency_contacts", c.APIBase, oid, eid), nil)

	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("page_index", ep.PageIndex)
	q.Add("item_per_page", ep.ItemPerPage)
	req.URL.RawQuery = q.Encode()
	response := &EmergencyContactListResponse{}

	err = c.SendWithAuth(req, response)
	return response, err
}
