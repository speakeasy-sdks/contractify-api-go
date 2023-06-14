# Contracts

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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Contracts.CreateContract(ctx, operations.CreateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                844266,
                602763,
                857946,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("illum"),
                },
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("vel"),
                },
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("error"),
                },
            },
            Departments: []int64{
                384382,
                437587,
                297534,
            },
            Documents: []int64{
                1,
                1,
                1,
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            IsOpenEnded: contractifyproduction.Bool(false),
            LegalEntities: []int64{
                963663,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                383441,
                477665,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                812169,
                528895,
                479977,
                568045,
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
        Company: 392785,
    }, operations.CreateContractSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.CreateContractRequest](../../models/operations/createcontractrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.CreateContractSecurity](../../models/operations/createcontractsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Contracts.DeleteContract(ctx, operations.DeleteContractRequest{
        Company: 925597,
        Contract: 836079,
    }, operations.DeleteContractSecurity{
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteContractRequest](../../models/operations/deletecontractrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteContractSecurity](../../models/operations/deletecontractsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Contracts.GetContract(ctx, operations.GetContractRequest{
        Company: 71036,
        Contract: 337396,
    }, operations.GetContractSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetContractRequest](../../models/operations/getcontractrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetContractSecurity](../../models/operations/getcontractsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Contracts.ListContracts(ctx, operations.ListContractsRequest{
        Company: 87129,
        Page: contractifyproduction.Int64(648172),
    }, operations.ListContractsSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListContractsRequest](../../models/operations/listcontractsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListContractsSecurity](../../models/operations/listcontractssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


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
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.Contracts.UpdateContract(ctx, operations.UpdateContractRequest{
        ContractWrite: &shared.ContractWrite{
            ContractTypes: []int64{
                368241,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("sapiente"),
                },
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("quo"),
                },
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("odit"),
                },
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("at"),
                },
            },
            Departments: []int64{
                978619,
                473608,
                799159,
                800911,
            },
            Documents: []int64{
                1,
                1,
            },
            DossierID: 1,
            Duration: contractifyproduction.String("P1Y"),
            EndDate: types.MustDateFromString("2021-12-31"),
            IsOpenEnded: contractifyproduction.Bool(false),
            LegalEntities: []int64{
                780529,
                678880,
                118274,
            },
            Name: "Partnership agreement",
            Offices: []int64{
                639921,
                582020,
                143353,
            },
            OwnerID: contractifyproduction.Int64(1),
            Phase: shared.ContractPhaseOngoing.ToPointer(),
            Relations: []int64{
                944669,
                758616,
                521848,
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
        Company: 105907,
        Contract: 414662,
    }, operations.UpdateContractSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateContractRequest](../../models/operations/updatecontractrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateContractSecurity](../../models/operations/updatecontractsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateContractResponse](../../models/operations/updatecontractresponse.md), error**

