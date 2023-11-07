# Offices
(*.Offices*)

### Available Operations

* [CreateOffice](#createoffice) - Create an office
* [DeleteOffice](#deleteoffice) - Delete an office
* [GetOffice](#getoffice) - Get an office
* [ListOffices](#listoffices) - List offices
* [UpdateOffice](#updateoffice) - Update an office

## CreateOffice

Create an office

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
    res, err := s.Offices.CreateOffice(ctx, operations.CreateOfficeRequest{
        OfficeWrite: &shared.OfficeWrite{
            Bus: contractifyproduction.String("1"),
            City: contractifyproduction.String("Sleidinge"),
            ContactPerson: contractifyproduction.String("Ada Lovelace"),
            Country: contractifyproduction.String("Belgium"),
            Email: contractifyproduction.String("info@contractify.be"),
            ID: contractifyproduction.Int64(1),
            Name: contractifyproduction.String("Ghent"),
            NumberIdentity: contractifyproduction.String("OFF-GHENT"),
            Phone: contractifyproduction.String("+32 9 234 28 97"),
            Street: contractifyproduction.String("Polenstraat 163"),
            Zip: contractifyproduction.String("9940"),
        },
        Company: 244393,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredAndOneApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateOfficeRequest](../../models/operations/createofficerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateOfficeResponse](../../models/operations/createofficeresponse.md), error**


## DeleteOffice

Delete an office

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
    res, err := s.Offices.DeleteOffice(ctx, operations.DeleteOfficeRequest{
        Company: 327183,
        Office: 668605,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.DeleteOfficeRequest](../../models/operations/deleteofficerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteOfficeResponse](../../models/operations/deleteofficeresponse.md), error**


## GetOffice

Get information about an office

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
    res, err := s.Offices.GetOffice(ctx, operations.GetOfficeRequest{
        Company: 616050,
        Office: 134885,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetOfficeRequest](../../models/operations/getofficerequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetOfficeResponse](../../models/operations/getofficeresponse.md), error**


## ListOffices

List all the offices within a company

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
    res, err := s.Offices.ListOffices(ctx, operations.ListOfficesRequest{
        Company: 303557,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OfficeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListOfficesRequest](../../models/operations/listofficesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListOfficesResponse](../../models/operations/listofficesresponse.md), error**


## UpdateOffice

Update an office

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
    res, err := s.Offices.UpdateOffice(ctx, operations.UpdateOfficeRequest{
        OfficeWrite: &shared.OfficeWrite{
            Bus: contractifyproduction.String("1"),
            City: contractifyproduction.String("Sleidinge"),
            ContactPerson: contractifyproduction.String("Ada Lovelace"),
            Country: contractifyproduction.String("Belgium"),
            Email: contractifyproduction.String("info@contractify.be"),
            ID: contractifyproduction.Int64(1),
            Name: contractifyproduction.String("Ghent"),
            NumberIdentity: contractifyproduction.String("OFF-GHENT"),
            Phone: contractifyproduction.String("+32 9 234 28 97"),
            Street: contractifyproduction.String("Polenstraat 163"),
            Zip: contractifyproduction.String("9940"),
        },
        Company: 989026,
        Office: 647378,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateOfficeRequest](../../models/operations/updateofficerequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateOfficeResponse](../../models/operations/updateofficeresponse.md), error**

