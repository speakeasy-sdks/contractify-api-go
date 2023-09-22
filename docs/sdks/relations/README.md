# Relations

### Available Operations

* [CreateRelation](#createrelation) - Create a relation
* [DeleteRelation](#deleterelation) - Delete a relation
* [GetRelation](#getrelation) - Get a relation
* [ListRelations](#listrelations) - List relations
* [UpdateRelation](#updaterelation) - Update a relation

## CreateRelation

Create a relation

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
    res, err := s.Relations.CreateRelation(ctx, operations.CreateRelationRequest{
        RelationWrite: &shared.RelationWrite{
            Address: &shared.Address{
                AddressLine1: contractifyproduction.String("221B Baker Street"),
                AddressLine2: contractifyproduction.String("Marylebone"),
                City: contractifyproduction.String("London"),
                Country: contractifyproduction.String("United Kingdom"),
                PostalCode: contractifyproduction.String("NW1 6XE"),
            },
            Email: contractifyproduction.String("sherlock@example.org"),
            Fax: contractifyproduction.String("+3211324354"),
            MobilePhone: contractifyproduction.String("+23477123456"),
            Name: "Sherlock Holmes Detective Services",
            Phone: contractifyproduction.String("+23477123456"),
            Reference: contractifyproduction.String("REF123"),
            Vat: contractifyproduction.String("BE12345678"),
            Website: contractifyproduction.String("https://www.example.org"),
        },
        Company: 105907,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateRelation201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateRelationRequest](../../models/operations/createrelationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateRelationResponse](../../models/operations/createrelationresponse.md), error**


## DeleteRelation

Delete a relation

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
    res, err := s.Relations.DeleteRelation(ctx, operations.DeleteRelationRequest{
        Company: 414662,
        Relation: 473600,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteRelationRequest](../../models/operations/deleterelationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteRelationResponse](../../models/operations/deleterelationresponse.md), error**


## GetRelation

Get information about a relation

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
    res, err := s.Relations.GetRelation(ctx, operations.GetRelationRequest{
        Company: 264555,
        Relation: 186332,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRelation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetRelationRequest](../../models/operations/getrelationrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetRelationResponse](../../models/operations/getrelationresponse.md), error**


## ListRelations

List all the relations within a company

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
    res, err := s.Relations.ListRelations(ctx, operations.ListRelationsRequest{
        Company: 774234,
        Page: contractifyproduction.Int64(736918),
        Reference: contractifyproduction.String("esse"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RelationCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListRelationsRequest](../../models/operations/listrelationsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListRelationsResponse](../../models/operations/listrelationsresponse.md), error**


## UpdateRelation

Update a relation

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
    res, err := s.Relations.UpdateRelation(ctx, operations.UpdateRelationRequest{
        RelationWrite: &shared.RelationWrite{
            Address: &shared.Address{
                AddressLine1: contractifyproduction.String("221B Baker Street"),
                AddressLine2: contractifyproduction.String("Marylebone"),
                City: contractifyproduction.String("London"),
                Country: contractifyproduction.String("United Kingdom"),
                PostalCode: contractifyproduction.String("NW1 6XE"),
            },
            Email: contractifyproduction.String("sherlock@example.org"),
            Fax: contractifyproduction.String("+3211324354"),
            MobilePhone: contractifyproduction.String("+23477123456"),
            Name: "Sherlock Holmes Detective Services",
            Phone: contractifyproduction.String("+23477123456"),
            Reference: contractifyproduction.String("REF123"),
            Vat: contractifyproduction.String("BE12345678"),
            Website: contractifyproduction.String("https://www.example.org"),
        },
        Company: 216550,
        Relation: 568434,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateRelation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateRelationRequest](../../models/operations/updaterelationrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateRelationResponse](../../models/operations/updaterelationresponse.md), error**

