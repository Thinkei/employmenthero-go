[![CircleCI](https://dl.circleci.com/status-badge/img/gh/Thinkei/employmenthero-go/tree/main.svg?style=svg&circle-token=d250e09ef4000c2bbd8e827fff9d1b36b6dd6cd4)](https://dl.circleci.com/status-badge/redirect/gh/Thinkei/employmenthero-go/tree/main) [![Coverage Status](https://coveralls.io/repos/github/Thinkei/employmenthero-go/badge.svg?branch=main)](https://coveralls.io/github/Thinkei/employmenthero-go?branch=main) [![Go Reference](https://pkg.go.dev/badge/github.com/Thinkei/employmenthero-go.svg)](https://pkg.go.dev/github.com/Thinkei/employmenthero-go) [![Go Report Card](https://goreportcard.com/badge/github.com/Thinkei/employmenthero-go)](https://goreportcard.com/report/github.com/Thinkei/employmenthero-go)

# Go client for EmploymentHero REST API
The official EmploymentHero Go client library.

## Requirements

- Go 1.16 or later

## Installation

Make sure your project is using Go Modules (it will have a `go.mod` file in its
root if it already is):

``` sh
go mod init
```

Then, reference stripe-go in a Go program with `import`:

``` go
import (
	"github.com/Thinkei/employmenthero-go"
)
```

Run any of the normal `go` commands (`build`/`install`/`test`). The Go
toolchain will resolve and fetch the stripe-go module automatically.

Alternatively, you can also explicitly `go get` the package into a project:

```bash
go get -u github.com/Thinkei/employmenthero-go
```

## Documentation

For a conprehensive list of examples, check out the [API documentation](https://developer.employmenthero.com/api-references/)

For details on all funtionality in this library, check out the [Go documentation](https://pkg.go.dev/github.com/Thinkei/employmenthero-go)

Below are a few of simple examples:

### Auth

```go
import "github.com/Thinkei/employmenthero-go"

// Create a client instance
c, err := employmenthero.NewClient("clientID", "secretID", "redirectUri", "OAuthHost", "apiHost")
c.SetLog(os.Stdout) // Set log to terminal stdout

// Get Authorization code and then use it to get the EH OAuth2 Access tokens
authroizationCode = "<authorizationCode>"
responseToken, err := client.GetOAuth2Access(ctx, authorizationCode)

// Save the refresh_token to anywhere you want,
// but use it when you call other APIs to get our resources
refreshToken := responseToken.RefreshToken
c.SetRefreshToken(responseToken.RefreshToken)

// call other APIs
organisationsResp, err := client.ListOrganisations(ctx, employmenthero.ListParams{})

if err != nil {
	fmt.Printf("Get Organisation failed - %s", err)
}

fmt.Println(organisationsResp.Data.Items)
```

### List Organisations

```go
response, err := c.ListOrganisations(context.TODO(), ListParams{})
organisations := response.Data.Items
```

### Get Organisation details

```go
response, err := c.GetOrganisation(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx")
organisation := response.Data
```

### List Employees

```go
response, err := c.ListEmployees(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
employees := response.Data.Items
```

### Get Employee details

```go
response, err := c.GetEmployee(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-yyyy")
employee := response.Data
```

### List Leave Requests

```go
response, err := c.ListLeaveRequests(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
leaveRequests := response.Data.Items
```

### Get Leave Request details

```go
response, err := c.GetLeaveRequest(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-yyyy")
leaveRequest := response.Data
```

### List Timesheet Entries

```go
response, err := c.ListTimesheetEntries(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "-", ListParams{})
timesheetEntries := response.Data.Items
```

### List Employemnt Histories

```go
response, err := c.ListEmploymentHistories(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "xxxxxx-yyyy", ListParams{})
employmentHistories := response.Data.Items
```

### List Emergency Contacts

```go
response, err := c.ListEmergencyContacts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "xxxx-yyy", ListParams{})
contacts := response.Data.Items
```

### List Teams

```go
response, err := c.ListTeams(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
teams := response.Data.Items
```

### List Employees by Team

```go
response, err := c.ListEmployeesByTeam(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXXX-YYYY-ZZZZ", ListParams{})
employees := response.Data.Items
```

### List Bank Accounts

```go
response, err := c.ListBankAccounts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
bankAccounts := response.Data.Items
```

### Get Tax Declaration of 1 Employee

```go
response, err := c.GetTaxDeclaration(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
taxDeclaration := response.Data
```

### Get Superannuation Detail of 1 Employee

```go
response, err := c.GetSuperannuationDetail(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
superannuationDetail := response.Data
```

### List Pay Details

```go
response, err := c.ListPayDetails(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
payDetails := response.Data.Items
```

### List Certification

```go
response, err := c.ListCertifications(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
certifications := response.Data.Items
```

### List Policies

```go
response, err := c.ListPolicies(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
policies := response.Data.Items
```

### List Employee Certifications

```go
response, err := c.ListEmployeeCertifications(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
employeeCertifications := response.Data.Items
```

### List Payslips

```go
response, err := c.ListPayslips(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
payslips := response.Data.Items
```

## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code must be `go fmt` compliant.
3. Ensure that `go test` succeeds.

## Test

The test suite needs testify's `require` package to run:

    github.com/stretchr/testify/require

Before running the tests, make sure to grab all of the package's dependencies:

    go get -t -v

Run all tests:

    make test

For any requests, bug or comments, please [open an issue][issues] or [submit a pull request][pulls].
