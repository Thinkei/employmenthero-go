# Go client for EmploymentHero REST API

## Usage

```go
import "github.com/Thinkei/employmenthero-go"

// Create a client instance
c, err := employmenthero.NewClient("clientID", "secretID", "refreshToken", "apiHost")

accessToken, err := c.GetAccessToken(context.Background())
```
