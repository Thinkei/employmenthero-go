package employmenthero

import (
	"net/http"
	"time"
)

const (
	RequestNewTokenBeforeExpiresIn = time.Duration(60) * time.Second
)

type (
	Client struct {
		ClientID       string
		Secret         string
		RedirectURI    string
		APIBase        string
		OAuthBase      string
		Client         HTTPClient
		Token          *TokenResponse
		tokenExpiresAt time.Time
	}

	TokenResponse struct {
		RefreshToken string `json:"refresh_token"`
		Token        string `json:"access_token"`
		Type         string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}

	ErrorResponse struct {
		Response *http.Response
	}

	HTTPClient interface {
		Do(req *http.Request) (*http.Response, error)
	}

	ListParams struct {
		PageIndex   string `json:"page_index"`    // Default: 1
		ItemPerPage string `json:"item_per_page"` // Default: 20
	}

	ListResponse struct {
		PageIndex   int `json:"page_index"`
		ItemPerPage int `json:"item_per_page"`
		TotalPages  int `json:"total_pages"`
		TotalItems  int `json:"total_items"`
	}

	Organisation struct {
		Id      string `json:"id"`
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Country string `json:"country"`
		LogoURL string `json:"logo_url"`
	}

	OrganisationDetail struct {
		Id                    string   `json:"id"`
		Name                  string   `json:"name"`
		Phone                 string   `json:"phone"`
		Country               string   `json:"country"`
		LogoURL               string   `json:"logo_url"`
		PrimaryAddress        string   `json:"primary_address"`
		EndOfWeek             string   `json:"end_of_week"`
		TypicalWorkDay        string   `json:"typical_work_day"`
		PayrollAdminEmails    []string `json:"payroll_admin_emails"`
		SubscriptionPlan      string   `json:"subscription_plan"`
		SuperfundName         string   `json:"superfund_name"`
		EmployeesCount        int      `json:"employees_count"`
		ActiveEmployeesCount  int      `json:"active_employees_count"`
		PendingEmployeesCount int      `json:"pending_employees_count"`
		TimeZone              string   `json:"time_zone"`
		CreatedAt             string   `json:"created_at"`
	}
)
