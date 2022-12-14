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

	BasicData struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	}

	Employee struct {
		Id                   string      `json:"id"`
		AccountEmail         string      `json:"account_email"`
		Title                string      `json:"title"`
		Role                 string      `json:"role"`
		FirstName            string      `json:"first_name"`
		LastName             string      `json:"last_name"`
		MiddleName           string      `json:"MiddleName"`
		Address              string      `json:"address"`
		AvatarURL            string      `json:"avatar_url"`
		KnownAs              string      `json:"known_as"`
		JobTitle             string      `json:"job_title"`
		Gender               string      `json:"gender"`
		Country              string      `json:"country"`
		Nationality          string      `json:"nationality"`
		DateOfBirth          string      `json:"date_of_birth"`
		MartialStatus        string      `json:"martial_status"`
		PersonalEmail        string      `json:"personal_email"`
		PersonalMobileNumber string      `json:"personal_mobile_number"`
		HomePhone            string      `json:"home_phone"`
		EmployingEntity      string      `json:"employing_entity"`
		Code                 string      `json:"code"`
		Location             string      `json:"location"`
		CompanyEmail         string      `json:"company_email"`
		CompanyNumber        string      `json:"company_number"`
		CompanyLandline      string      `json:"company_landline"`
		StartDate            string      `json:"start_date"`
		TerminationDate      string      `json:"termination_date"`
		PrimaryCostCentre    BasicData   `json:"primary_cost_centre"`
		SecondaryCostCentres []BasicData `json:"secondary_cost_centres"`
		PrimaryManager       BasicData   `json:"primary_manager"`
		SecondaryManager     BasicData   `json:"secondary_manager"`
		ExternalId           string      `json:"external_id"`
	}

	LeaveRequest struct {
		Id                 string  `json:"id"`
		StartDate          string  `json:"start_date"`
		EndDate            string  `json:"end_date"`
		TotalHours         float32 `json:"total_hours"`
		Comment            string  `json:"comment"`
		Status             string  `json:"status"`
		LeaveBalanceAmount float32 `json:"leave_balance_amount"`
		LeaveCategoryName  string  `json:"leave_category_name"`
		Reason             string  `json:"reason"`
		EmployeeId         string  `json:"employee_id"`
	}

	Timesheet struct {
		Id         string    `json:"id"`
		EmployeeId string    `json:"employee_id"`
		Date       string    `json:"date"`
		StartTime  string    `json:"start_time"`
		EndTime    string    `json:"end_time"`
		Status     string    `json:"status"`
		Units      float32   `json:"units"`
		Reason     string    `json:"reason"`
		Comment    string    `json:"comment"`
		Time       int       `json:"time"`
		CostCentre BasicData `json:"cost_centre"`
	}

	EmploymentHistory struct {
		Id             string `json:"id"`
		Title          string `json:"title"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		EmploymentType string `json:"employment_type"`
	}

	EmergencyContact struct {
		Id                   int    `json:"id"`
		ContactName          string `json:"contact_name"`
		DaytimeContactNumber string `json:"daytime_contact_number"`
		AfterHoursNo         string `json:"after_hours_no"`
		AfterMobileNo        string `json:"after_mobile_no"`
		Relationship         string `json:"relationship"`
		ContactType          string `json:"contact_type"`
	}

	Team struct {
		Id     string `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	BankAccount struct {
		Id             string `json:"id"`
		AccountName    string `json:"account_name"`
		AccountNumber  string `json:"account_number"`
		Bsb            string `json:"bsb"`
		Amount         string `json:"amount"`
		PrimaryAccount bool   `json:"primary_account"`
	}

	TaxDeclaration struct {
		FirstName                  string `json:"first_name"`
		LastName                   string `json:"last_name"`
		TaxFileNumber              string `json:"tax_file_number"`
		TaxAuResident              bool   `json:"tax_au_resident"`
		TaxForeignResident         bool   `json:"tax_foreign_resident"`
		WorkingHolidayMaker        bool   `json:"working_holiday_maker"`
		TaxFree                    bool   `json:"tax_free"`
		TaxHelpDebt                bool   `json:"tax_help_debt"`
		TaxFinancialSupplementDebt bool   `json:"tax_financial_supplement_debt"`
	}

	SuperannuationDetail struct {
		FundName                 string `json:"fund_name"`
		MemberNumber             string `json:"member_number"`
		ProductCode              string `json:"product_code"`
		EmployerNominatedFund    bool   `json:"employer_nominated_fund"`
		FundAbn                  string `json:"fund_abn"`
		ElectronicServiceAddress string `json:"electronic_service_address"`
		FundEmail                string `json:"fund_email"`
		AccountName              string `json:"account_name"`
		AccountBSB               string `json:"account_bsb"`
		AccountNumber            string `json:"account_number"`
	}

	PayDetail struct {
		Id                     string  `json:"id"`
		EffectiveFrom          string  `json:"effective_from"`
		Classification         string  `json:"classification"`
		IndustrialInstrument   string  `json:"industrial_instrument"`
		PayRateTemplate        string  `json:"pay_rate_template"`
		AnniversaryDate        string  `json:"anniversary_date"`
		Salary                 float32 `json:"salary"`
		SalaryType             string  `json:"salary_type"`
		PayUnit                string  `json:"pay_unit"`
		PayCategory            string  `json:"pay_category"`
		LeaveAllowanceTemplate string  `json:"leave_allowance_template"`
		ChangeReason           string  `json:"change_reason"`
		Comments               string  `json:"comments"`
	}

	Certification struct {
		Id   string `json:"id"`
		Name string `json:"name"`
		Type string `json:"type"`
	}

	Policy struct {
		Id        string `json:"id"`
		Name      string `json:"name"`
		Induction bool   `json:"induction"`
		CreatedAt string `json:"created_at"`
	}
)
