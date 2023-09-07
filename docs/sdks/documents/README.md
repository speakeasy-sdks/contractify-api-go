# Documents

### Available Operations

* [DeleteDocument](#deletedocument) - Delete a document
* [GetDocument](#getdocument) - Get a document
* [ListDocuments](#listdocuments) - List documents
* [UpdateDocument](#updatedocument) - Update a document

## DeleteDocument

Delete a document

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
    operationSecurity := operations.DeleteDocumentSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Documents.DeleteDocument(ctx, operations.DeleteDocumentRequest{
        Company: 20218,
        Document: 368241,
    }, operationSecurity)
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
| `request`                                                                              | [operations.DeleteDocumentRequest](../../models/operations/deletedocumentrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.DeleteDocumentSecurity](../../models/operations/deletedocumentsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.DeleteDocumentResponse](../../models/operations/deletedocumentresponse.md), error**


## GetDocument

Get information about a document

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
    operationSecurity := operations.GetDocumentSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Documents.GetDocument(ctx, operations.GetDocumentRequest{
        Company: 832620,
        Document: 957156,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDocument200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.GetDocumentRequest](../../models/operations/getdocumentrequest.md)   | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `security`                                                                       | [operations.GetDocumentSecurity](../../models/operations/getdocumentsecurity.md) | :heavy_check_mark:                                                               | The security requirements to use for the request.                                |


### Response

**[*operations.GetDocumentResponse](../../models/operations/getdocumentresponse.md), error**


## ListDocuments

List all the documents within a company

### Example Usage

```go
package main

import(
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.ListDocumentsSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Documents.ListDocuments(ctx, operations.ListDocumentsRequest{
        Company: 778157,
        EsigningStatus: operations.ListDocumentsEsigningStatusSentToLegal.ToPointer(),
        EsigningUpdatedAfter: types.MustTimeFromString("2020-05-23T06:06:25.221Z"),
        Page: contractifyproduction.Int64(978619),
        RelationID: contractifyproduction.Int64(473608),
        SignedAfter: types.MustTimeFromString("2020-08-07T00:03:55.328Z"),
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.DocumentCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDocumentsRequest](../../models/operations/listdocumentsrequest.md)   | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |
| `security`                                                                           | [operations.ListDocumentsSecurity](../../models/operations/listdocumentssecurity.md) | :heavy_check_mark:                                                                   | The security requirements to use for the request.                                    |


### Response

**[*operations.ListDocumentsResponse](../../models/operations/listdocumentsresponse.md), error**


## UpdateDocument

Update a document

### Example Usage

```go
package main

import(
	"context"
	"log"
	"ContractifyProduction"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/models/shared"
)

func main() {
    s := contractifyproduction.New()
    operationSecurity := operations.UpdateDocumentSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.Documents.UpdateDocument(ctx, operations.UpdateDocumentRequest{
        DocumentWrite: &shared.DocumentWrite{
            Contracts: []int64{
                1,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("esse"),
                },
            },
            Description: contractifyproduction.String("Lorem ipsum dolor sit amet."),
            Dossiers: []int64{
                1,
            },
            Name: "filename.pdf",
            OwnerID: contractifyproduction.Int64(1),
        },
        Company: 520478,
        Document: 780529,
    }, operationSecurity)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDocument200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateDocumentRequest](../../models/operations/updatedocumentrequest.md)   | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |
| `security`                                                                             | [operations.UpdateDocumentSecurity](../../models/operations/updatedocumentsecurity.md) | :heavy_check_mark:                                                                     | The security requirements to use for the request.                                      |


### Response

**[*operations.UpdateDocumentResponse](../../models/operations/updatedocumentresponse.md), error**

