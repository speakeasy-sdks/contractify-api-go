# Contracts
(*Contracts*)

### Available Operations

* [CreateContract](#createcontract) - Create a contract
* [DeleteContract](#deletecontract) - Delete a contract
* [GetContract](#getcontract) - Get a contract
* [ListContracts](#listcontracts) - List contracts
* [UpdateContract](#updatecontract) - Update a contract

## CreateContract

Create a contract

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
    res, err := s.Contracts.CreateContract(ctx, operations.CreateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                1,
                2,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("software"),
                },
            },
            Departments: []int64{
                1,
                2,
            },
            Documents: []int64{
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            LegalEntities: []int64{
                1,
                2,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                1,
                2,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                1,
                2,
            },
            Renewal: &shared.ContractRenewal{
                AutomaticRenewal: &shared.ContractAutomaticRenewal{
                    NumberOfRenewals: contractifyproduction.Int64(1),
                    RenewalPeriod: contractifyproduction.String("P3Y"),
                },
            },
            StartDate: types.MustDateFromString("2021-01-01"),
            Termination: &shared.ContractTermination{
                TerminationDate: types.MustDateFromString("2021-11-30"),
                TerminationDuration: contractifyproduction.String("P1M"),
            },
        },
        Company: 940947,
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
| `request`                                                                                | [operations.CreateContractRequest](../../pkg/models/operations/createcontractrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateContractResponse](../../pkg/models/operations/createcontractresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.CreateContractResponseBody                  | 401                                                   | application/json                                      |
| sdkerrors.CreateContractContractsResponseBody         | 403                                                   | application/json                                      |
| sdkerrors.CreateContractContractsResponseResponseBody | 422                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |

## DeleteContract

Delete a contract

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
    res, err := s.Contracts.DeleteContract(ctx, operations.DeleteContractRequest{
        Company: 791005,
        Contract: 200974,
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
| `request`                                                                                | [operations.DeleteContractRequest](../../pkg/models/operations/deletecontractrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteContractResponse](../../pkg/models/operations/deletecontractresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.DeleteContractResponseBody                     | 400                                                      | application/json                                         |
| sdkerrors.DeleteContractContractsResponseBody            | 401                                                      | application/json                                         |
| sdkerrors.DeleteContractContractsResponseResponseBody    | 403                                                      | application/json                                         |
| sdkerrors.DeleteContractContractsResponse404ResponseBody | 404                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 400-600                                                  | */*                                                      |

## GetContract

Get information about a contract

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
    res, err := s.Contracts.GetContract(ctx, operations.GetContractRequest{
        Company: 362495,
        Contract: 163842,
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
| `request`                                                                          | [operations.GetContractRequest](../../pkg/models/operations/getcontractrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetContractResponse](../../pkg/models/operations/getcontractresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.GetContractResponseBody                  | 401                                                | application/json                                   |
| sdkerrors.GetContractContractsResponseBody         | 403                                                | application/json                                   |
| sdkerrors.GetContractContractsResponseResponseBody | 404                                                | application/json                                   |
| sdkerrors.SDKError                                 | 400-600                                            | */*                                                |

## ListContracts

List all the contracts within a company

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
    res, err := s.Contracts.ListContracts(ctx, operations.ListContractsRequest{
        Company: 567515,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContractCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListContractsRequest](../../pkg/models/operations/listcontractsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListContractsResponse](../../pkg/models/operations/listcontractsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.ListContractsResponseBody          | 401                                          | application/json                             |
| sdkerrors.ListContractsContractsResponseBody | 403                                          | application/json                             |
| sdkerrors.SDKError                           | 400-600                                      | */*                                          |

## UpdateContract

Update a contract

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
    res, err := s.Contracts.UpdateContract(ctx, operations.UpdateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                1,
                2,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("software"),
                },
            },
            Departments: []int64{
                1,
                2,
            },
            Documents: []int64{
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            LegalEntities: []int64{
                1,
                2,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                1,
                2,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                1,
                2,
            },
            Renewal: &shared.ContractRenewal{
                AutomaticRenewal: &shared.ContractAutomaticRenewal{
                    NumberOfRenewals: contractifyproduction.Int64(1),
                    RenewalPeriod: contractifyproduction.String("P3Y"),
                },
            },
            StartDate: types.MustDateFromString("2021-01-01"),
            Termination: &shared.ContractTermination{
                TerminationDate: types.MustDateFromString("2021-11-30"),
                TerminationDuration: contractifyproduction.String("P1M"),
            },
        },
        Company: 60280,
        Contract: 331790,
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
| `request`                                                                                | [operations.UpdateContractRequest](../../pkg/models/operations/updatecontractrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateContractResponse](../../pkg/models/operations/updatecontractresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.UpdateContractResponseBody                     | 401                                                      | application/json                                         |
| sdkerrors.UpdateContractContractsResponseBody            | 403                                                      | application/json                                         |
| sdkerrors.UpdateContractContractsResponseResponseBody    | 404                                                      | application/json                                         |
| sdkerrors.UpdateContractContractsResponse422ResponseBody | 422                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 400-600                                                  | */*                                                      |
