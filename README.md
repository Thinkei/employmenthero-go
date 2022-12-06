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

## Get Organisations

```go
response, err := c.ListOrganisations(context.TODO(), OrganisationListParams{})
organistaions := response.Data.Items
```

## How to Contribute

- Create an issue to describe what you want to do
- Fork a repository
- Add/Fix something
- Check that tests are passing
- Create PR

## Tests

- Unit tests: go test -v ./...


