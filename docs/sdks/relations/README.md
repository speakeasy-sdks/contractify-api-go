# Relations
(*Relations*)

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
        Company: 528070,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateRelationRequest](../../pkg/models/operations/createrelationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateRelationResponse](../../pkg/models/operations/createrelationresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.CreateRelationResponseBody                  | 401                                                   | application/json                                      |
| sdkerrors.CreateRelationRelationsResponseBody         | 403                                                   | application/json                                      |
| sdkerrors.CreateRelationRelationsResponseResponseBody | 422                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 4xx-5xx                                               | */*                                                   |

## DeleteRelation

Delete a relation

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Relations.DeleteRelation(ctx, operations.DeleteRelationRequest{
        Company: 773418,
        Relation: 890630,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteRelationRequest](../../pkg/models/operations/deleterelationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteRelationResponse](../../pkg/models/operations/deleterelationresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.DeleteRelationResponseBody                     | 400                                                      | application/json                                         |
| sdkerrors.DeleteRelationRelationsResponseBody            | 401                                                      | application/json                                         |
| sdkerrors.DeleteRelationRelationsResponseResponseBody    | 403                                                      | application/json                                         |
| sdkerrors.DeleteRelationRelationsResponse404ResponseBody | 404                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4xx-5xx                                                  | */*                                                      |

## GetRelation

Get information about a relation

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
    res, err := s.Relations.GetRelation(ctx, operations.GetRelationRequest{
        Company: 734058,
        Relation: 979643,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetRelationRequest](../../pkg/models/operations/getrelationrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetRelationResponse](../../pkg/models/operations/getrelationresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.GetRelationResponseBody                  | 401                                                | application/json                                   |
| sdkerrors.GetRelationRelationsResponseBody         | 403                                                | application/json                                   |
| sdkerrors.GetRelationRelationsResponseResponseBody | 404                                                | application/json                                   |
| sdkerrors.SDKError                                 | 4xx-5xx                                            | */*                                                |

## ListRelations

List all the relations within a company

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
    res, err := s.Relations.ListRelations(ctx, operations.ListRelationsRequest{
        Company: 454135,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListRelationsRequest](../../pkg/models/operations/listrelationsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListRelationsResponse](../../pkg/models/operations/listrelationsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.ListRelationsResponseBody          | 401                                          | application/json                             |
| sdkerrors.ListRelationsRelationsResponseBody | 403                                          | application/json                             |
| sdkerrors.SDKError                           | 4xx-5xx                                      | */*                                          |

## UpdateRelation

Update a relation

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
        Company: 573397,
        Relation: 281147,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateRelationRequest](../../pkg/models/operations/updaterelationrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateRelationResponse](../../pkg/models/operations/updaterelationresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.UpdateRelationResponseBody                     | 401                                                      | application/json                                         |
| sdkerrors.UpdateRelationRelationsResponseBody            | 403                                                      | application/json                                         |
| sdkerrors.UpdateRelationRelationsResponseResponseBody    | 404                                                      | application/json                                         |
| sdkerrors.UpdateRelationRelationsResponse422ResponseBody | 422                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 4xx-5xx                                                  | */*                                                      |
