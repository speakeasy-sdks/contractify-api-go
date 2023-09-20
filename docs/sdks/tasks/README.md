# Tasks

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
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := ContractifyProduction.New()
    operationSecurity := operations.CreateTaskSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Tasks.CreateTask(ctx, operations.CreateTaskRequest{
        TaskWrite: &shared.TaskWrite{
            ContractID: ContractifyProduction.Int64(1),
            Description: ContractifyProduction.String("Lorem ipsum dolor sit amet."),
            DueDate: types.MustDateFromString("2021-12-31"),
            DueDateDependsOn: shared.TaskWriteDueDateDependsOnEndDate.ToPointer(),
            DueDateInterval: ContractifyProduction.String("-P10D"),
            OwnerID: ContractifyProduction.Int64(1),
            ReminderDuration: ContractifyProduction.String("P1M"),
            RepetitionInterval: ContractifyProduction.String("P1Y"),
            Status: shared.TaskWriteStatusAccomplished.ToPointer(),
            Title: ContractifyProduction.String("My task"),
        },
        Company: 135218,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateTaskRequest](../../models/operations/createtaskrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.CreateTaskSecurity](../../models/operations/createtasksecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.CreateTaskResponse](../../models/operations/createtaskresponse.md), error**


## DeleteTask

Delete a task

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
    s := ContractifyProduction.New()
    operationSecurity := operations.DeleteTaskSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Tasks.DeleteTask(ctx, operations.DeleteTaskRequest{
        Company: 18789,
        Task: 324141,
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.DeleteTaskRequest](../../models/operations/deletetaskrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.DeleteTaskSecurity](../../models/operations/deletetasksecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.DeleteTaskResponse](../../models/operations/deletetaskresponse.md), error**


## GetTask

Get a task

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
    s := ContractifyProduction.New()
    operationSecurity := operations.GetTaskSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Tasks.GetTask(ctx, operations.GetTaskRequest{
        Company: 617636,
        Task: 149675,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `request`                                                                | [operations.GetTaskRequest](../../models/operations/gettaskrequest.md)   | :heavy_check_mark:                                                       | The request object to use for the request.                               |
| `security`                                                               | [operations.GetTaskSecurity](../../models/operations/gettasksecurity.md) | :heavy_check_mark:                                                       | The security requirements to use for the request.                        |


### Response

**[*operations.GetTaskResponse](../../models/operations/gettaskresponse.md), error**


## ListTasks

List all tasks within a company

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
    s := ContractifyProduction.New()
    operationSecurity := operations.ListTasksSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Tasks.ListTasks(ctx, operations.ListTasksRequest{
        Company: 612096,
        Page: ContractifyProduction.Int64(222321),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.TaskCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListTasksRequest](../../models/operations/listtasksrequest.md)   | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `security`                                                                   | [operations.ListTasksSecurity](../../models/operations/listtaskssecurity.md) | :heavy_check_mark:                                                           | The security requirements to use for the request.                            |


### Response

**[*operations.ListTasksResponse](../../models/operations/listtasksresponse.md), error**


## UpdateTask

Update a task

### Example Usage

```go
package main

import(
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := ContractifyProduction.New()
    operationSecurity := operations.UpdateTaskSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Tasks.UpdateTask(ctx, operations.UpdateTaskRequest{
        TaskUpdate: &shared.TaskUpdate{
            Description: ContractifyProduction.String("Lorem ipsum dolor sit amet."),
            DueDate: types.MustDateFromString("2021-12-31"),
            DueDateDependsOn: shared.TaskUpdateDueDateDependsOnEndDate.ToPointer(),
            DueDateInterval: ContractifyProduction.String("-P10D"),
            OwnerID: ContractifyProduction.Int64(1),
            ReminderDuration: ContractifyProduction.String("P1M"),
            RepetitionInterval: ContractifyProduction.String("P1Y"),
            Status: shared.TaskUpdateStatusAccomplished.ToPointer(),
            Title: ContractifyProduction.String("My task"),
        },
        Company: 616934,
        Task: 386489,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateTask200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.UpdateTaskRequest](../../models/operations/updatetaskrequest.md)   | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `security`                                                                     | [operations.UpdateTaskSecurity](../../models/operations/updatetasksecurity.md) | :heavy_check_mark:                                                             | The security requirements to use for the request.                              |


### Response

**[*operations.UpdateTaskResponse](../../models/operations/updatetaskresponse.md), error**

