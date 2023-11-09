# Departments
(*Departments*)

### Available Operations

* [CreateDepartment](#createdepartment) - Create a department
* [DeleteDepartment](#deletedepartment) - Delete a department
* [GetDepartment](#getdepartment) - Get a department
* [ListDepartments](#listdepartments) - List departments
* [UpdateDepartment](#updatedepartment) - Update a department

## CreateDepartment

Create a department

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
    res, err := s.Departments.CreateDepartment(ctx, operations.CreateDepartmentRequest{
        DepartmentWrite: &shared.DepartmentWrite{
            Name: contractifyproduction.String("Sales"),
        },
        Company: 33324,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.CreateDepartmentRequest](../../pkg/models/operations/createdepartmentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.CreateDepartmentResponse](../../pkg/models/operations/createdepartmentresponse.md), error**
| Error Object                                              | Status Code                                               | Content Type                                              |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| sdkerrors.CreateDepartmentResponseBody                    | 401                                                       | application/json                                          |
| sdkerrors.CreateDepartmentDepartmentsResponseBody         | 403                                                       | application/json                                          |
| sdkerrors.CreateDepartmentDepartmentsResponseResponseBody | 422                                                       | application/json                                          |
| sdkerrors.SDKError                                        | 400-600                                                   | */*                                                       |

## DeleteDepartment

Delete a department

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
    res, err := s.Departments.DeleteDepartment(ctx, operations.DeleteDepartmentRequest{
        Company: 701942,
        Department: 751163,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.DeleteDepartmentRequest](../../pkg/models/operations/deletedepartmentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.DeleteDepartmentResponse](../../pkg/models/operations/deletedepartmentresponse.md), error**
| Error Object                                                 | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| sdkerrors.DeleteDepartmentResponseBody                       | 400                                                          | application/json                                             |
| sdkerrors.DeleteDepartmentDepartmentsResponseBody            | 401                                                          | application/json                                             |
| sdkerrors.DeleteDepartmentDepartmentsResponseResponseBody    | 403                                                          | application/json                                             |
| sdkerrors.DeleteDepartmentDepartmentsResponse404ResponseBody | 404                                                          | application/json                                             |
| sdkerrors.SDKError                                           | 400-600                                                      | */*                                                          |

## GetDepartment

Get information about a department

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
    res, err := s.Departments.GetDepartment(ctx, operations.GetDepartmentRequest{
        Company: 255130,
        Department: 855529,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetDepartmentRequest](../../pkg/models/operations/getdepartmentrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetDepartmentResponse](../../pkg/models/operations/getdepartmentresponse.md), error**
| Error Object                                           | Status Code                                            | Content Type                                           |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| sdkerrors.GetDepartmentResponseBody                    | 401                                                    | application/json                                       |
| sdkerrors.GetDepartmentDepartmentsResponseBody         | 403                                                    | application/json                                       |
| sdkerrors.GetDepartmentDepartmentsResponseResponseBody | 404                                                    | application/json                                       |
| sdkerrors.SDKError                                     | 400-600                                                | */*                                                    |

## ListDepartments

List all the departments within a company

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
    res, err := s.Departments.ListDepartments(ctx, operations.ListDepartmentsRequest{
        Company: 117069,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DepartmentCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListDepartmentsRequest](../../pkg/models/operations/listdepartmentsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.ListDepartmentsResponse](../../pkg/models/operations/listdepartmentsresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.ListDepartmentsResponseBody            | 401                                              | application/json                                 |
| sdkerrors.ListDepartmentsDepartmentsResponseBody | 403                                              | application/json                                 |
| sdkerrors.SDKError                               | 400-600                                          | */*                                              |

## UpdateDepartment

Update a department

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
    res, err := s.Departments.UpdateDepartment(ctx, operations.UpdateDepartmentRequest{
        DepartmentWrite: &shared.DepartmentWrite{
            Name: contractifyproduction.String("Sales"),
        },
        Company: 431122,
        Department: 2342,
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateDepartmentRequest](../../pkg/models/operations/updatedepartmentrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateDepartmentResponse](../../pkg/models/operations/updatedepartmentresponse.md), error**
| Error Object                                                 | Status Code                                                  | Content Type                                                 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| sdkerrors.UpdateDepartmentResponseBody                       | 401                                                          | application/json                                             |
| sdkerrors.UpdateDepartmentDepartmentsResponseBody            | 403                                                          | application/json                                             |
| sdkerrors.UpdateDepartmentDepartmentsResponseResponseBody    | 404                                                          | application/json                                             |
| sdkerrors.UpdateDepartmentDepartmentsResponse422ResponseBody | 422                                                          | application/json                                             |
| sdkerrors.SDKError                                           | 400-600                                                      | */*                                                          |
