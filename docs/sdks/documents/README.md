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
    res, err := s.Documents.DeleteDocument(ctx, operations.DeleteDocumentRequest{
        Company: 368241,
        Document: 832620,
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.DeleteDocumentRequest](../../models/operations/deletedocumentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Documents.GetDocument(ctx, operations.GetDocumentRequest{
        Company: 957156,
        Document: 778157,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetDocument200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetDocumentRequest](../../models/operations/getdocumentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
	contractifyproduction "ContractifyProduction"
	"ContractifyProduction/pkg/models/shared"
	"ContractifyProduction/pkg/models/operations"
	"ContractifyProduction/pkg/types"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: "",
            PersonalAccessToken: "",
        }),
    )

    ctx := context.Background()
    res, err := s.Documents.ListDocuments(ctx, operations.ListDocumentsRequest{
        Company: 140350,
        EsigningStatus: operations.ListDocumentsEsigningStatusFinishedButPartiallySigned.ToPointer(),
        EsigningUpdatedAfter: types.MustTimeFromString("2020-01-25T09:54:35.794Z"),
        Page: contractifyproduction.Int64(473608),
        RelationID: contractifyproduction.Int64(799159),
        SignedAfter: types.MustTimeFromString("2021-08-13T16:19:19.906Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DocumentCollection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListDocumentsRequest](../../models/operations/listdocumentsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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
    res, err := s.Documents.UpdateDocument(ctx, operations.UpdateDocumentRequest{
        DocumentWrite: &shared.DocumentWrite{
            Contracts: []int64{
                1,
            },
            CustomFieldValues: []shared.CustomFieldValueWrite{
                shared.CustomFieldValueWrite{
                    CustomFieldID: contractifyproduction.Int64(2),
                    Value: contractifyproduction.String("totam"),
                },
            },
            Description: contractifyproduction.String("Lorem ipsum dolor sit amet."),
            Dossiers: []int64{
                1,
            },
            Name: "filename.pdf",
            OwnerID: contractifyproduction.Int64(1),
        },
        Company: 780529,
        Document: 678880,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateDocument200ApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.UpdateDocumentRequest](../../models/operations/updatedocumentrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.UpdateDocumentResponse](../../models/operations/updatedocumentresponse.md), error**

