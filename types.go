package employmenthero

import (
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	RequestNewTokenBeforeExpiresIn = time.Duration(60) * time.Second // The TTL (Time To Live) of the access token

)

type (
	// A Client is an entry point of this package, it includes all needded configurations
	// and functions that we need to call EmploymentHero Apis
	Client struct {
		// A mutex is used to prevent race condition when re-new or assign token to Client instance
		mu sync.Mutex
		// ClientID is a unique string representing the registration OAuth 2.0 application.
		ClientID string
		// Secret is a random secret used by client to authenticate to the Employment Hero Authorisation Server.
		Secret string
		// RedirectURI is one of your specified redirect url(s) in your OAuth 2.0 application
		RedirectURI string
		// The host of RESTful EmploymentHero API - https://api.employmenthero.com
		APIBase string
		// The host of OAuth2 EmploymentHero API - https://oauth.employmenthero.com
		OAuthBase      string
		Client         HTTPClient
		Token          *TokenResponse
		Log            io.Writer
		tokenExpiresAt time.Time
	}

	// A response object sent from /oauth/token API which
	TokenResponse struct {
		RefreshToken string `json:"refresh_token"`
		Token        string `json:"access_token"`
		Type         string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
	}

	// Reason for the failure.
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

	// Organisation is an object of one Organisation resource in EH system.
	// It must be managed by you or the organisation which you work for.
	Organisation struct {
		Id      string `json:"id"` // Unique identifier for the object. (UUID format)
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Country string `json:"country"`
		LogoURL string `json:"logo_url"`
	}

	// OrganisationDetail is the details of one Organisation resource in EH system.
	// It must be managed by you or the organisation which you work for.
	// [Organisation Example]: https://developer.employmenthero.com/api-references/#the-organisation-object
	OrganisationDetail struct {
		Id                    string   `json:"id"` // Unique identifier for the object. (UUID format)
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

	// Employee is the details of one Employee resource in EH system.
	// [Employee Example]: https://developer.employmenthero.com/api-references/#the-employee-object
	Employee struct {
		Id                   string      `json:"id"` // Unique identifier for the object. (UUID format)
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

	// LeaveRequest is the details of one LeaveRequest resource in EH system.
	// [LeaveRequest Example]: https://developer.employmenthero.com/api-references/#the-leave-request-object
	LeaveRequest struct {
		Id                 string  `json:"id"` // Unique identifier for the object. (UUID format)
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

	// Timesheet is the details of one TimesheetEntry resource in EH system.
	// [TimesheetEntry Example]: https://developer.employmenthero.com/api-references/#the-timesheet-object
	Timesheet struct {
		Id         string    `json:"id"` // Unique identifier for the object. (UUID format)
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

	// EmploymentHistory is the details of one employment history resource in EH system.
	// [EmploymentHistory Example]: https://developer.employmenthero.com/api-references/#the-employment-history-object
	EmploymentHistory struct {
		Id             string `json:"id"` // Unique identifier for the object. (UUID format)
		Title          string `json:"title"`
		StartDate      string `json:"start_date"`
		EndDate        string `json:"end_date"`
		EmploymentType string `json:"employment_type"`
	}

	// EmergencyContact is the details of one EmergencyContact resource in EH system.
	// [EmergencyContact Example]: https://developer.employmenthero.com/api-references/#the-emergency-contact-object
	EmergencyContact struct {
		Id                   int    `json:"id"` // Unique identifier for the object. (Integer)
		ContactName          string `json:"contact_name"`
		DaytimeContactNumber string `json:"daytime_contact_number"`
		AfterHoursNo         string `json:"after_hours_no"`
		AfterMobileNo        string `json:"after_mobile_no"`
		Relationship         string `json:"relationship"`
		ContactType          string `json:"contact_type"`
	}

	// Team is the details of one Team resource in EH system.
	// [Team Example]: https://developer.employmenthero.com/api-references/#the-team-object
	Team struct {
		Id     string `json:"id"` // Unique identifier for the object. (UUID format)
		Name   string `json:"name"`
		Status string `json:"status"`
	}

	// BankAccount is the details of one BankAccount resource in EH system.
	// [BankAccount Example]: https://developer.employmenthero.com/api-references/#the-bank-account-object
	BankAccount struct {
		Id             string `json:"id"` // Unique identifier for the object. (UUID format)
		AccountName    string `json:"account_name"`
		AccountNumber  string `json:"account_number"`
		Bsb            string `json:"bsb"`
		Amount         string `json:"amount"`
		PrimaryAccount bool   `json:"primary_account"`
	}

	// TaxDeclaration is the details of one TaxDeclaration resource in EH system.
	// [TaxDeclaration Example]: https://developer.employmenthero.com/api-references/#the-tax-declaration-object
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

	// SuperannuationDetail is the details of one SuperannuationDetail resource in EH system.
	// [SuperannuationDetail Example]: https://developer.employmenthero.com/api-references/#the-superannuation-detail-object
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

	// PayDetail is the details of one PayDetail resource in EH system.
	// [PayDetail Example]: https://developer.employmenthero.com/api-references/#the-pay-detail-object
	PayDetail struct {
		Id                     string  `json:"id"` // Unique identifier for the object. (UUID format)
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

	// Certification is the details of one Certification resource in EH system.
	// [Certification Example]: https://developer.employmenthero.com/api-references/#the-certification-object
	Certification struct {
		Id   string `json:"id"` // Unique identifier for the object. (UUID format)
		Name string `json:"name"`
		Type string `json:"type"`
	}

	// EmployeeCertification is the details of one Certification resource of one employee in EH system
	// [EmployeeCertification Example]: https://developer.employmenthero.com/api-references/#the-employee-certification-object
	EmployeeCertification struct {
		Id              string `json:"id"` // Unique identifier for the object. (UUID format)
		Name            string `json:"name"`
		CertificationId string `json:"certification_id"`
		Type            string `json:"type"`
		ExpiryDate      string `json:"expiry_date"`
		CompletionDate  string `json:"completion_date"`
		DriverProblem   bool   `json:"driver_problem"`
		DriverDetails   string `json:"driver_details"`
		Status          string `json:"status"`
		Reason          string `json:"reason"`
	}

	// Policy is the details of one Policy resource in EH system
	// [Policy Example]: https://developer.employmenthero.com/api-references/#get-policies
	Policy struct {
		Id        string `json:"id"` // Unique identifier for the object. (UUID format)
		Name      string `json:"name"`
		Induction bool   `json:"induction"`
		CreatedAt string `json:"created_at"`
	}

	// Payslip is the details of one Payslip resource in EH system
	// [Payslip Example]: https://developer.employmenthero.com/api-references/#the-payslip-object
	Payslip struct {
		Id               string  `json:"id"` // Unique identifier for the object. (UUID format)
		FirstName        string  `json:"first_name"`
		LastName         string  `json:"last_name"`
		TotalDeduction   int     `json:"total_deduction"`
		NetPay           float32 `json:"net_pay"`
		Super            float32 `json:"super"`
		Wages            float32 `json:"wages"`
		Reimbursements   float32 `json:"reimbursements"`
		Tax              float32 `json:"tax"`
		Name             string  `json:"name"`
		AddressLine1     string  `json:"address_line1"`
		AddressLine2     string  `json:"address_line2"`
		Suburb           string  `json:"suburb"`
		Postcode         string  `json:"postcode"`
		PostTaxDeduction float32 `json:"post_tax_deduction"`
		PreTaxDeduction  float32 `json:"pre_tax_deduction"`
		BaseRate         float32 `json:"base_rate"`
		HourlyRate       float32 `json:"hourly_rate"`
		BaseRateUnit     string  `json:"base_rate_unit"`
		EmploymentType   string  `json:"employment_type"`
		PaymentDate      string  `json:"payment_date"`
	}
)
