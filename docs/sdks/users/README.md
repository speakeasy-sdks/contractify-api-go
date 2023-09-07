# Users

### Available Operations

* [CurrentUser](#currentuser) - Current User
* [ListUsers](#listusers) - List users

## CurrentUser

Get the current user details

### Example Usage

```go
package main

import(
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.CurrentUserSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Users.CurrentUser(ctx, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CurrentUser200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `security`                                                                       | [operations.CurrentUserSecurity](../../models/operations/currentusersecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.CurrentUserResponse](../../models/operations/currentuserresponse.md), error**


## ListUsers

List all the users within a company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.ListUsersSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Users.ListUsers(ctx, operations.ListUsersRequest{
        Company: 943749,
        Page: contractifyproduction.Int64(902599),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UserCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListUsersRequest](../../models/operations/listusersrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.ListUsersSecurity](../../models/operations/listuserssecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.ListUsersResponse](../../models/operations/listusersresponse.md), error**

