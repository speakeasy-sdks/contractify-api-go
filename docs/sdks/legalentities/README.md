# LegalEntities
(*LegalEntities*)

### Available Operations

* [ListLegalEntities](#listlegalentities) - List legal entities

## ListLegalEntities

List all the legal entities within a company

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
    res, err := s.LegalEntities.ListLegalEntities(ctx, operations.ListLegalEntitiesRequest{
        Company: 730248,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LegalEntityCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListLegalEntitiesRequest](../../pkg/models/operations/listlegalentitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.ListLegalEntitiesResponse](../../pkg/models/operations/listlegalentitiesresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.ListLegalEntitiesResponseBody              | 401                                                  | application/json                                     |
| sdkerrors.ListLegalEntitiesLegalEntitiesResponseBody | 403                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |
