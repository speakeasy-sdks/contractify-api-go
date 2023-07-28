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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.CreateRelationSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

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
        Company: 607831,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateRelation201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateRelationRequest](../../models/operations/createrelationrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateRelationSecurity](../../models/operations/createrelationsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.DeleteRelationSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Relations.DeleteRelation(ctx, operations.DeleteRelationRequest{
        Company: 363711,
        Relation: 325047,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteRelationRequest](../../models/operations/deleterelationrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteRelationSecurity](../../models/operations/deleterelationsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.GetRelationSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Relations.GetRelation(ctx, operations.GetRelationRequest{
        Company: 570197,
        Relation: 38425,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRelation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetRelationRequest](../../models/operations/getrelationrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetRelationSecurity](../../models/operations/getrelationsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.ListRelationsSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Relations.ListRelations(ctx, operations.ListRelationsRequest{
        Company: 438601,
        Page: contractifyproduction.Int64(634274),
        Reference: contractifyproduction.String("doloribus"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.RelationCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListRelationsRequest](../../models/operations/listrelationsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListRelationsSecurity](../../models/operations/listrelationssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.UpdateRelationSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

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
        Company: 958950,
        Relation: 102044,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateRelation200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateRelationRequest](../../models/operations/updaterelationrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateRelationSecurity](../../models/operations/updaterelationsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateRelationResponse](../../models/operations/updaterelationresponse.md), error**

