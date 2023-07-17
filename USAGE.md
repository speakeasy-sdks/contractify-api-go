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
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
        Company: 548814,
    }, operations.ListContractTypesSecurity{
        OAuth2: "",
        PersonalAccessToken: "",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ContractTypeCollection != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->