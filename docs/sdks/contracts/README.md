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
	"context"
	"log"
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: "",
            PersonalAccessToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Contracts.CreateContract(ctx, operations.CreateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                940947,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("Palladium"),
                },
            },
            Departments: []int64{
                559447,
            },
            Documents: []int64{
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            IsOpenEnded: contractifyproduction.Bool(false),
            LegalEntities: []int64{
                657681,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                410255,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                867240,
            },
            Renewal: &shared.ContractRenewal{
                AutomaticRenewal: &shared.ContractAutomaticRenewal{
                    NumberOfRenewals: contractifyproduction.Int64(1),
                    RenewalPeriod: contractifyproduction.String("P3Y"),
                },
                IsAutomaticallyRenewed: contractifyproduction.Bool(false),
            },
            StartDate: types.MustDateFromString("2021-01-01"),
            Termination: &shared.ContractTermination{
                IsTerminableAtAnyTime: contractifyproduction.Bool(false),
                TerminationDate: types.MustDateFromString("2021-11-30"),
                TerminationDuration: contractifyproduction.String("P1M"),
            },
        },
        Company: 647814,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateContract201ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateContractRequest](../../models/operations/createcontractrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.CreateContractResponse](../../models/operations/createcontractresponse.md), error**


## DeleteContract

Delete a contract

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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteContractRequest](../../models/operations/deletecontractrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.DeleteContractResponse](../../models/operations/deletecontractresponse.md), error**


## GetContract

Get information about a contract

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
    res, err := s.Contracts.GetContract(ctx, operations.GetContractRequest{
        Company: 362495,
        Contract: 163842,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetContractRequest](../../models/operations/getcontractrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetContractResponse](../../models/operations/getcontractresponse.md), error**


## ListContracts

List all the contracts within a company

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
    res, err := s.Contracts.ListContracts(ctx, operations.ListContractsRequest{
        Company: 567515,
        Page: contractifyproduction.Int64(62076),
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

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListContractsRequest](../../models/operations/listcontractsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListContractsResponse](../../models/operations/listcontractsresponse.md), error**


## UpdateContract

Update a contract

### Example Usage

```go
package main

import(
	"context"
	"log"
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: "",
            PersonalAccessToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Contracts.UpdateContract(ctx, operations.UpdateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                60280,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("synergies"),
                },
            },
            Departments: []int64{
                178046,
            },
            Documents: []int64{
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            IsOpenEnded: contractifyproduction.Bool(false),
            LegalEntities: []int64{
                331597,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                693387,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                670029,
            },
            Renewal: &shared.ContractRenewal{
                AutomaticRenewal: &shared.ContractAutomaticRenewal{
                    NumberOfRenewals: contractifyproduction.Int64(1),
                    RenewalPeriod: contractifyproduction.String("P3Y"),
                },
                IsAutomaticallyRenewed: contractifyproduction.Bool(false),
            },
            StartDate: types.MustDateFromString("2021-01-01"),
            Termination: &shared.ContractTermination{
                IsTerminableAtAnyTime: contractifyproduction.Bool(false),
                TerminationDate: types.MustDateFromString("2021-11-30"),
                TerminationDuration: contractifyproduction.String("P1M"),
            },
        },
        Company: 213877,
        Contract: 177820,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateContract200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateContractRequest](../../models/operations/updatecontractrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateContractResponse](../../models/operations/updatecontractresponse.md), error**

