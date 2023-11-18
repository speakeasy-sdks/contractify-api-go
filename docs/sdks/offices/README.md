# Offices
(*Offices*)

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
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
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

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateOfficeRequest](../../pkg/models/operations/createofficerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateOfficeResponse](../../pkg/models/operations/createofficeresponse.md), error**
| Error Object                                      | Status Code                                       | Content Type                                      |
| ------------------------------------------------- | ------------------------------------------------- | ------------------------------------------------- |
| sdkerrors.CreateOfficeResponseBody                | 401                                               | application/json                                  |
| sdkerrors.CreateOfficeOfficesResponseBody         | 403                                               | application/json                                  |
| sdkerrors.CreateOfficeOfficesResponseResponseBody | 422                                               | application/json                                  |
| sdkerrors.SDKError                                | 400-600                                           | */*                                               |

## DeleteOffice

Delete an office

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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteOfficeRequest](../../pkg/models/operations/deleteofficerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteOfficeResponse](../../pkg/models/operations/deleteofficeresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.DeleteOfficeResponseBody                   | 400                                                  | application/json                                     |
| sdkerrors.DeleteOfficeOfficesResponseBody            | 401                                                  | application/json                                     |
| sdkerrors.DeleteOfficeOfficesResponseResponseBody    | 403                                                  | application/json                                     |
| sdkerrors.DeleteOfficeOfficesResponse404ResponseBody | 404                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |

## GetOffice

Get information about an office

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

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetOfficeRequest](../../pkg/models/operations/getofficerequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetOfficeResponse](../../pkg/models/operations/getofficeresponse.md), error**
| Error Object                                   | Status Code                                    | Content Type                                   |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| sdkerrors.GetOfficeResponseBody                | 401                                            | application/json                               |
| sdkerrors.GetOfficeOfficesResponseBody         | 403                                            | application/json                               |
| sdkerrors.GetOfficeOfficesResponseResponseBody | 404                                            | application/json                               |
| sdkerrors.SDKError                             | 400-600                                        | */*                                            |

## ListOffices

List all the offices within a company

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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListOfficesRequest](../../pkg/models/operations/listofficesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListOfficesResponse](../../pkg/models/operations/listofficesresponse.md), error**
| Error Object                             | Status Code                              | Content Type                             |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| sdkerrors.ListOfficesResponseBody        | 401                                      | application/json                         |
| sdkerrors.ListOfficesOfficesResponseBody | 403                                      | application/json                         |
| sdkerrors.SDKError                       | 400-600                                  | */*                                      |

## UpdateOffice

Update an office

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

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateOfficeRequest](../../pkg/models/operations/updateofficerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateOfficeResponse](../../pkg/models/operations/updateofficeresponse.md), error**
| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.UpdateOfficeResponseBody                   | 401                                                  | application/json                                     |
| sdkerrors.UpdateOfficeOfficesResponseBody            | 403                                                  | application/json                                     |
| sdkerrors.UpdateOfficeOfficesResponseResponseBody    | 404                                                  | application/json                                     |
| sdkerrors.UpdateOfficeOfficesResponse422ResponseBody | 422                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |
