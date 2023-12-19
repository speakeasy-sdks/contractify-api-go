# Tasks
(*Tasks*)

### Available Operations

* [CreateTask](#createtask) - Create a task
* [DeleteTask](#deletetask) - Delete a task
* [GetTask](#gettask) - Get a task
* [ListTasks](#listtasks) - List tasks
* [UpdateTask](#updatetask) - Update a task

## CreateTask

Create a task

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/types"
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
    res, err := s.Tasks.CreateTask(ctx, operations.CreateTaskRequest{
        TaskWrite: &shared.TaskWrite{
            ContractID: contractifyproduction.Int64(1),
            Description: contractifyproduction.String("Lorem ipsum dolor sit amet."),
            DueDate: types.MustDateFromString("2021-12-31"),
            DueDateDependsOn: shared.TaskWriteDueDateDependsOnEndDate.ToPointer(),
            DueDateInterval: contractifyproduction.String("-P10D"),
            OwnerID: contractifyproduction.Int64(1),
            ReminderDuration: contractifyproduction.String("P1M"),
            RepetitionInterval: contractifyproduction.String("P1Y"),
            Status: shared.TaskWriteStatusAccomplished.ToPointer(),
            Title: contractifyproduction.String("My task"),
        },
        Company: 296904,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.CreateTaskRequest](../../pkg/models/operations/createtaskrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreateTaskResponse](../../pkg/models/operations/createtaskresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.CreateTaskResponseBody              | 401                                           | application/json                              |
| sdkerrors.CreateTaskTasksResponseBody         | 403                                           | application/json                              |
| sdkerrors.CreateTaskTasksResponseResponseBody | 422                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## DeleteTask

Delete a task

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
    res, err := s.Tasks.DeleteTask(ctx, operations.DeleteTaskRequest{
        Company: 357574,
        Task: 394977,
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
| `request`                                                                        | [operations.DeleteTaskRequest](../../pkg/models/operations/deletetaskrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.DeleteTaskResponse](../../pkg/models/operations/deletetaskresponse.md), error**
| Error Object                                  | Status Code                                   | Content Type                                  |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| sdkerrors.DeleteTaskResponseBody              | 401                                           | application/json                              |
| sdkerrors.DeleteTaskTasksResponseBody         | 403                                           | application/json                              |
| sdkerrors.DeleteTaskTasksResponseResponseBody | 404                                           | application/json                              |
| sdkerrors.SDKError                            | 4xx-5xx                                       | */*                                           |

## GetTask

Get a task

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
    res, err := s.Tasks.GetTask(ctx, operations.GetTaskRequest{
        Company: 717011,
        Task: 649018,
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

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetTaskRequest](../../pkg/models/operations/gettaskrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.GetTaskResponse](../../pkg/models/operations/gettaskresponse.md), error**
| Error Object                               | Status Code                                | Content Type                               |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| sdkerrors.GetTaskResponseBody              | 401                                        | application/json                           |
| sdkerrors.GetTaskTasksResponseBody         | 403                                        | application/json                           |
| sdkerrors.GetTaskTasksResponseResponseBody | 404                                        | application/json                           |
| sdkerrors.SDKError                         | 4xx-5xx                                    | */*                                        |

## ListTasks

List all tasks within a company

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
    res, err := s.Tasks.ListTasks(ctx, operations.ListTasksRequest{
        Company: 715197,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListTasksRequest](../../pkg/models/operations/listtasksrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListTasksResponse](../../pkg/models/operations/listtasksresponse.md), error**
| Error Object                         | Status Code                          | Content Type                         |
| ------------------------------------ | ------------------------------------ | ------------------------------------ |
| sdkerrors.ListTasksResponseBody      | 401                                  | application/json                     |
| sdkerrors.ListTasksTasksResponseBody | 403                                  | application/json                     |
| sdkerrors.SDKError                   | 4xx-5xx                              | */*                                  |

## UpdateTask

Update a task

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/types"
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
    res, err := s.Tasks.UpdateTask(ctx, operations.UpdateTaskRequest{
        TaskUpdate: &shared.TaskUpdate{
            Description: contractifyproduction.String("Lorem ipsum dolor sit amet."),
            DueDate: types.MustDateFromString("2021-12-31"),
            DueDateDependsOn: shared.TaskUpdateDueDateDependsOnEndDate.ToPointer(),
            DueDateInterval: contractifyproduction.String("-P10D"),
            OwnerID: contractifyproduction.Int64(1),
            ReminderDuration: contractifyproduction.String("P1M"),
            RepetitionInterval: contractifyproduction.String("P1Y"),
            Status: shared.TaskUpdateStatusAccomplished.ToPointer(),
            Title: contractifyproduction.String("My task"),
        },
        Company: 449699,
        Task: 675064,
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.UpdateTaskRequest](../../pkg/models/operations/updatetaskrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.UpdateTaskResponse](../../pkg/models/operations/updatetaskresponse.md), error**
| Error Object                                     | Status Code                                      | Content Type                                     |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| sdkerrors.UpdateTaskResponseBody                 | 401                                              | application/json                                 |
| sdkerrors.UpdateTaskTasksResponseBody            | 403                                              | application/json                                 |
| sdkerrors.UpdateTaskTasksResponseResponseBody    | 404                                              | application/json                                 |
| sdkerrors.UpdateTaskTasksResponse422ResponseBody | 422                                              | application/json                                 |
| sdkerrors.SDKError                               | 4xx-5xx                                          | */*                                              |
