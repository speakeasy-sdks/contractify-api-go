<!-- Start SDK Example Usage [usage] -->
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
			OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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
<!-- End SDK Example Usage [usage] -->