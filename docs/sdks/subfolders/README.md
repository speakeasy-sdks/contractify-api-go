# Subfolders
(*Subfolders*)

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
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: "",
            PersonalAccessToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Subfolders.ListSubfolders(ctx, operations.ListSubfoldersRequest{
        Company: 749068,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DossierCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListSubfoldersRequest](../../models/operations/listsubfoldersrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListSubfoldersResponse](../../models/operations/listsubfoldersresponse.md), error**

