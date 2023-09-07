# Offices

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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.CreateOfficeSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

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
        Company: 118274,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateOffice201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.CreateOfficeRequest](../../models/operations/createofficerequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.CreateOfficeSecurity](../../models/operations/createofficesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.DeleteOfficeSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Offices.DeleteOffice(ctx, operations.DeleteOfficeRequest{
        Company: 720633,
        Office: 639921,
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.DeleteOfficeRequest](../../models/operations/deleteofficerequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.DeleteOfficeSecurity](../../models/operations/deleteofficesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.GetOfficeSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Offices.GetOffice(ctx, operations.GetOfficeRequest{
        Company: 582020,
        Office: 143353,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetOffice200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetOfficeRequest](../../models/operations/getofficerequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.GetOfficeSecurity](../../models/operations/getofficesecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.ListOfficesSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Offices.ListOffices(ctx, operations.ListOfficesRequest{
        Company: 537373,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.OfficeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListOfficesRequest](../../models/operations/listofficesrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.ListOfficesSecurity](../../models/operations/listofficessecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.UpdateOfficeSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

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
        Company: 944669,
        Office: 758616,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateOffice200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.UpdateOfficeRequest](../../models/operations/updateofficerequest.md)   | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `security`                                                                         | [operations.UpdateOfficeSecurity](../../models/operations/updateofficesecurity.md) | :heavy_check_mark:                                                                 | The security requirements to use for the request.                                  |


### Response

**[*operations.UpdateOfficeResponse](../../models/operations/updateofficeresponse.md), error**

