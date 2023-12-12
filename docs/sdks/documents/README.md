# Documents
(*Documents*)

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
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
	"net/http"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Documents.DeleteDocument(ctx, operations.DeleteDocumentRequest{
        Company: 431806,
        Document: 379848,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteDocumentRequest](../../pkg/models/operations/deletedocumentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteDocumentResponse](../../pkg/models/operations/deletedocumentresponse.md), error**
| Error Object                                             | Status Code                                              | Content Type                                             |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| sdkerrors.DeleteDocumentResponseBody                     | 401                                                      | application/json                                         |
| sdkerrors.DeleteDocumentDocumentsResponseBody            | 403                                                      | application/json                                         |
| sdkerrors.DeleteDocumentDocumentsResponseResponseBody    | 404                                                      | application/json                                         |
| sdkerrors.DeleteDocumentDocumentsResponse422ResponseBody | 422                                                      | application/json                                         |
| sdkerrors.SDKError                                       | 400-600                                                  | */*                                                      |

## GetDocument

Get information about a document

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Documents.GetDocument(ctx, operations.GetDocumentRequest{
        Company: 67810,
        Document: 267106,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetDocumentRequest](../../pkg/models/operations/getdocumentrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetDocumentResponse](../../pkg/models/operations/getdocumentresponse.md), error**
| Error Object                                       | Status Code                                        | Content Type                                       |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| sdkerrors.GetDocumentResponseBody                  | 401                                                | application/json                                   |
| sdkerrors.GetDocumentDocumentsResponseBody         | 403                                                | application/json                                   |
| sdkerrors.GetDocumentDocumentsResponseResponseBody | 404                                                | application/json                                   |
| sdkerrors.SDKError                                 | 400-600                                            | */*                                                |

## ListDocuments

List all the documents within a company

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Documents.ListDocuments(ctx, operations.ListDocumentsRequest{
        Company: 581480,
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListDocumentsRequest](../../pkg/models/operations/listdocumentsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListDocumentsResponse](../../pkg/models/operations/listdocumentsresponse.md), error**
| Error Object                                 | Status Code                                  | Content Type                                 |
| -------------------------------------------- | -------------------------------------------- | -------------------------------------------- |
| sdkerrors.ListDocumentsResponseBody          | 401                                          | application/json                             |
| sdkerrors.ListDocumentsDocumentsResponseBody | 403                                          | application/json                             |
| sdkerrors.SDKError                           | 400-600                                      | */*                                          |

## UpdateDocument

Update a document

### Example Usage

```go
package main

import(
	"ContractifyProduction/pkg/models/shared"
	contractifyproduction "ContractifyProduction"
	"context"
	"ContractifyProduction/pkg/models/operations"
	"log"
)

func main() {
    s := contractifyproduction.New(
        contractifyproduction.WithSecurity(shared.Security{
            OAuth2: contractifyproduction.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
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
                    Value: contractifyproduction.String("software"),
                },
            },
            Description: contractifyproduction.String("Lorem ipsum dolor sit amet."),
            Dossiers: []int64{
                1,
            },
            Name: "filename.pdf",
            OwnerID: contractifyproduction.Int64(1),
        },
        Company: 653381,
        Document: 312704,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateDocumentRequest](../../pkg/models/operations/updatedocumentrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateDocumentResponse](../../pkg/models/operations/updatedocumentresponse.md), error**
| Error Object                                          | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| sdkerrors.UpdateDocumentResponseBody                  | 401                                                   | application/json                                      |
| sdkerrors.UpdateDocumentDocumentsResponseBody         | 403                                                   | application/json                                      |
| sdkerrors.UpdateDocumentDocumentsResponseResponseBody | 404                                                   | application/json                                      |
| sdkerrors.SDKError                                    | 400-600                                               | */*                                                   |
