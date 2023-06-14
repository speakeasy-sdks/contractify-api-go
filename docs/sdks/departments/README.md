# Departments

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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Departments.CreateDepartment(ctx, operations.CreateDepartmentRequest{
        DepartmentWrite: &shared.DepartmentWrite{
            Name: contractifyproduction.String("Sales"),
        },
        Company: 264555,
    }, operations.CreateDepartmentSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateDepartment201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.CreateDepartmentRequest](../../models/operations/createdepartmentrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.CreateDepartmentSecurity](../../models/operations/createdepartmentsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.CreateDepartmentResponse](../../models/operations/createdepartmentresponse.md), error**


## DeleteDepartment

Delete a department

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

    ctx := context.Background()
    res, err := s.Departments.DeleteDepartment(ctx, operations.DeleteDepartmentRequest{
        Company: 186332,
        Department: 774234,
    }, operations.DeleteDepartmentSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteDepartmentRequest](../../models/operations/deletedepartmentrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.DeleteDepartmentSecurity](../../models/operations/deletedepartmentsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.DeleteDepartmentResponse](../../models/operations/deletedepartmentresponse.md), error**


## GetDepartment

Get information about a department

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

    ctx := context.Background()
    res, err := s.Departments.GetDepartment(ctx, operations.GetDepartmentRequest{
        Company: 736918,
        Department: 456150,
    }, operations.GetDepartmentSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDepartment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetDepartmentRequest](../../models/operations/getdepartmentrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.GetDepartmentSecurity](../../models/operations/getdepartmentsecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.GetDepartmentResponse](../../models/operations/getdepartmentresponse.md), error**


## ListDepartments

List all the departments within a company

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

    ctx := context.Background()
    res, err := s.Departments.ListDepartments(ctx, operations.ListDepartmentsRequest{
        Company: 216550,
    }, operations.ListDepartmentsSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListDepartmentsRequest](../../models/operations/listdepartmentsrequest.md)   | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `security`                                                                               | [operations.ListDepartmentsSecurity](../../models/operations/listdepartmentssecurity.md) | :heavy_check_mark:                                                                       | The security requirements to use for the request.                                        |


### Response

**[*operations.ListDepartmentsResponse](../../models/operations/listdepartmentsresponse.md), error**


## UpdateDepartment

Update a department

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

    ctx := context.Background()
    res, err := s.Departments.UpdateDepartment(ctx, operations.UpdateDepartmentRequest{
        DepartmentWrite: &shared.DepartmentWrite{
            Name: contractifyproduction.String("Sales"),
        },
        Company: 568434,
        Department: 135218,
    }, operations.UpdateDepartmentSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDepartment200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.UpdateDepartmentRequest](../../models/operations/updatedepartmentrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.UpdateDepartmentSecurity](../../models/operations/updatedepartmentsecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.UpdateDepartmentResponse](../../models/operations/updatedepartmentresponse.md), error**

