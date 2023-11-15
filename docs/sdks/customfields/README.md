# CustomFields
(*CustomFields*)

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
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/models/operations"
)

func main() {
    s := contractifyproduction.New()

    ctx := context.Background()
    res, err := s.CustomFields.ListCustomFields(ctx, operations.ListCustomFieldsRequest{
        Company: 318971,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomFieldCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListCustomFieldsRequest](../../pkg/models/operations/listcustomfieldsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListCustomFieldsResponse](../../pkg/models/operations/listcustomfieldsresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.ListCustomFieldsResponseBody             | 401                                                | application/json                                   |
| sdkerrors.ListCustomFieldsCustomFieldsResponseBody | 403                                                | application/json                                   |
| sdkerrors.SDKError                                 | 400-600                                            | */*                                                |
