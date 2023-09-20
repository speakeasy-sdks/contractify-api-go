# CustomFields

### Available Operations

* [ListCustomFields](#listcustomfields) - List custom fields

## ListCustomFields

List all the custom fields within a company

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
    operationSecurity := operations.ListCustomFieldsSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.CustomFields.ListCustomFields(ctx, operations.ListCustomFieldsRequest{
        Company: 479977,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomFieldCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.ListCustomFieldsRequest](../../models/operations/listcustomfieldsrequest.md)   | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |
| `security`                                                                                 | [operations.ListCustomFieldsSecurity](../../models/operations/listcustomfieldssecurity.md) | :heavy_check_mark:                                                                         | The security requirements to use for the request.                                          |


### Response

**[*operations.ListCustomFieldsResponse](../../models/operations/listcustomfieldsresponse.md), error**

