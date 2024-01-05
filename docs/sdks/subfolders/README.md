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
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListSubfoldersRequest](../../pkg/models/operations/listsubfoldersrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListSubfoldersResponse](../../pkg/models/operations/listsubfoldersresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.ListSubfoldersResponseBody           | 401                                            | application/json                               |
| sdkerrors.ListSubfoldersSubfoldersResponseBody | 403                                            | application/json                               |
| sdkerrors.SDKError                             | 4xx-5xx                                        | */*                                            |
