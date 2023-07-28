# Subfolders

### Available Operations

* [ListSubfolders](#listsubfolders) - List subfolders

## ListSubfolders

List all the subfolders within a company

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
    operationSecurity := operations.ListSubfoldersSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Subfolders.ListSubfolders(ctx, operations.ListSubfoldersRequest{
        Company: 652790,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DossierCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListSubfoldersRequest](../../models/operations/listsubfoldersrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.ListSubfoldersSecurity](../../models/operations/listsubfolderssecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.ListSubfoldersResponse](../../models/operations/listsubfoldersresponse.md), error**

