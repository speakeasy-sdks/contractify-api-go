# ContractTypes
(*ContractTypes*)

### Available Operations

* [ListContractTypes](#listcontracttypes) - List contract types

## ListContractTypes

List all the contract types within a company

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
    res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
        Company: 839467,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContractTypeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListContractTypesRequest](../../pkg/models/operations/listcontracttypesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListContractTypesResponse](../../pkg/models/operations/listcontracttypesresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.ListContractTypesResponseBody              | 401                                                  | application/json                                     |
| sdkerrors.ListContractTypesContractTypesResponseBody | 403                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |
