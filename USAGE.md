<!-- Start SDK Example Usage -->


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
        Company: 548814,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.ContractTypeCollection != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->