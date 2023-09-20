# ContractTypes

### Available Operations

* [ListContractTypes](#listcontracttypes) - List contract types

## ListContractTypes

List all the contract types within a company

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
    operationSecurity := operations.ListContractTypesSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
        Company: 592845,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ContractTypeCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListContractTypesRequest](../../models/operations/listcontracttypesrequest.md)   | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `security`                                                                                   | [operations.ListContractTypesSecurity](../../models/operations/listcontracttypessecurity.md) | :heavy_check_mark:                                                                           | The security requirements to use for the request.                                            |


### Response

**[*operations.ListContractTypesResponse](../../models/operations/listcontracttypesresponse.md), error**

