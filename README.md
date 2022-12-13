![main workflow](https://github.com/Thinkei/employmenthero-go/actions/workflows/go.yml/badge.svg)

# Go client for EmploymentHero REST API
This repository contains a Go package for the client-side components of the EmploymentHero APIs

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

## How to Contribute

- Create an issue to describe what you want to do
- Fork a repository
- Add/Fix something
- Check that tests are passing
- Create PR

## Tests

- Unit tests: go test -v ./...


