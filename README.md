![main workflow](https://github.com/Thinkei/employmenthero-go/actions/workflows/go.yml/badge.svg)

# Go client for EmploymentHero REST API
The official EmploymentHero Go client library.

## Requirements

- Go 1.16 or later

## Usage

```go
import "github.com/Thinkei/employmenthero-go"

// Create a client instance
c, err := employmenthero.NewClient("clientID", "secretID", "refreshToken", "OAuthHost", "apiHost")

accessToken, err := c.GetAccessToken(context.TODO())
```

## List Organisations

```go
response, err := c.ListOrganisations(context.TODO(), ListParams{})
organistaions := response.Data.Items
```

## Get Organisation details

```go
response, err := c.GetOrganisation(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx")
organistaion := response.Data
```

## List Employees

```go
response, err := c.ListEmployees(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
employees := response.Data.Items
```

## Get Employee details

```go
response, err := c.GetEmployee(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-yyyy")
employee := response.Data
```

## List Leave Requests

```go
response, err := c.ListLeaveRequests(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
leaveRequests := response.Data.Items
```

## Get Leave Request details

```go
response, err := c.GetLeaveRequest(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "90a34ef1-50e4-4930-a9d6-yyyy")
leaveRequest := response.Data
```

## List Timesheet Entries

```go
response, err := c.ListTimesheetEntries(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "-", ListParams{})
timesheetEntries := response.Data.Items
```

## List Employemnt Histories

```go
response, err := c.ListEmploymentHistories(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "xxxxxx-yyyy", ListParams{})
employmentHistories := response.Data.Items
```

## List Emergency Contacts

```go
response, err := c.ListEmergencyContacts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "xxxx-yyy", ListParams{})
contacts := response.Data.Items
```

## List Teams

```go
response, err := c.ListTeams(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", ListParams{})
teams := response.Data.Items
```

## List Employees by Team

```go
response, err := c.ListEmployeesByTeam(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXXX-YYYY-ZZZZ", ListParams{})
employees := response.Data.Items
```

## List Bank Accounts

```go
response, err := c.ListBankAccounts(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
bankAccounts := response.Data.Items
```

## Get Tax Declaration of 1 Employee

```go
response, err := c.GetTaxDeclaration(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
taxDeclaration := response.Data
```

## Get Superannuation Detail of 1 Employee

```go
response, err := c.GetSuperannuationDetail(context.TODO(), "90a34ef1-50e4-4930-a9d6-xxxx", "XXX-YY-ZZZ", ListParams{})
superannuationDetail := response.Data
```

## Development

Pull requests from the community are welcome. If you submit one, please keep
the following guidelines in mind:

1. Code must be `go fmt` compliant.
3. Ensure that `go test` succeeds.

## Tests

- Unit tests: go test -v ./...


