<!-- Start SDK Example Usage -->
```go
package main

import (
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
	"context"
	"log"
)

func main() {
	s := contractifyproduction.New(
		contractifyproduction.WithSecurity(shared.Security{
			OAuth2:              "",
			PersonalAccessToken: "",
		}),
	)

	ctx := context.Background()
	res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
		Company: 839467,
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