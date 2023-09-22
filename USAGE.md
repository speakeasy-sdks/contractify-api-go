<!-- Start SDK Example Usage -->


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
    res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
        Company: 548814,
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