# LegalEntities

### Available Operations

* [ListLegalEntities](#listlegalentities) - List legal entities

## ListLegalEntities

List all the legal entities within a company

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
    res, err := s.LegalEntities.ListLegalEntities(ctx, operations.ListLegalEntitiesRequest{
        Company: 118274,
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListLegalEntitiesRequest](../../models/operations/listlegalentitiesrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListLegalEntitiesResponse](../../models/operations/listlegalentitiesresponse.md), error**

