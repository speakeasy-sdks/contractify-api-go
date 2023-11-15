# ContractifyProduction

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/contractify-api-go.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/contractify-api-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>


## 🏗 **Welcome to your new SDK!** 🏗

It has been generated successfully based on your OpenAPI spec. However, it is not yet ready for production use. Here are some next steps:
- [ ] 🛠 Make your SDK feel handcrafted by [customizing it](https://www.speakeasyapi.dev/docs/customize-sdks)
- [ ] ♻️ Refine your SDK quickly by iterating locally with the [Speakeasy CLI](https://github.com/speakeasy-api/speakeasy)
- [ ] 🎁 Publish your SDK to package managers by [configuring automatic publishing](https://www.speakeasyapi.dev/docs/productionize-sdks/publish-sdks)
- [ ] ✨ When ready to productionize, delete this section from the README
<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/contractify-api-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
### Example

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
	s := contractifyproduction.New()

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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [ContractTypes](docs/sdks/contracttypes/README.md)

* [ListContractTypes](docs/sdks/contracttypes/README.md#listcontracttypes) - List contract types

### [Contracts](docs/sdks/contracts/README.md)

* [CreateContract](docs/sdks/contracts/README.md#createcontract) - Create a contract
* [DeleteContract](docs/sdks/contracts/README.md#deletecontract) - Delete a contract
* [GetContract](docs/sdks/contracts/README.md#getcontract) - Get a contract
* [ListContracts](docs/sdks/contracts/README.md#listcontracts) - List contracts
* [UpdateContract](docs/sdks/contracts/README.md#updatecontract) - Update a contract

### [CustomFields](docs/sdks/customfields/README.md)

* [ListCustomFields](docs/sdks/customfields/README.md#listcustomfields) - List custom fields

### [Departments](docs/sdks/departments/README.md)

* [CreateDepartment](docs/sdks/departments/README.md#createdepartment) - Create a department
* [DeleteDepartment](docs/sdks/departments/README.md#deletedepartment) - Delete a department
* [GetDepartment](docs/sdks/departments/README.md#getdepartment) - Get a department
* [ListDepartments](docs/sdks/departments/README.md#listdepartments) - List departments
* [UpdateDepartment](docs/sdks/departments/README.md#updatedepartment) - Update a department

### [Documents](docs/sdks/documents/README.md)

* [DeleteDocument](docs/sdks/documents/README.md#deletedocument) - Delete a document
* [GetDocument](docs/sdks/documents/README.md#getdocument) - Get a document
* [ListDocuments](docs/sdks/documents/README.md#listdocuments) - List documents
* [UpdateDocument](docs/sdks/documents/README.md#updatedocument) - Update a document

### [Subfolders](docs/sdks/subfolders/README.md)

* [ListSubfolders](docs/sdks/subfolders/README.md#listsubfolders) - List subfolders

### [LegalEntities](docs/sdks/legalentities/README.md)

* [ListLegalEntities](docs/sdks/legalentities/README.md#listlegalentities) - List legal entities

### [Offices](docs/sdks/offices/README.md)

* [CreateOffice](docs/sdks/offices/README.md#createoffice) - Create an office
* [DeleteOffice](docs/sdks/offices/README.md#deleteoffice) - Delete an office
* [GetOffice](docs/sdks/offices/README.md#getoffice) - Get an office
* [ListOffices](docs/sdks/offices/README.md#listoffices) - List offices
* [UpdateOffice](docs/sdks/offices/README.md#updateoffice) - Update an office

### [Relations](docs/sdks/relations/README.md)

* [CreateRelation](docs/sdks/relations/README.md#createrelation) - Create a relation
* [DeleteRelation](docs/sdks/relations/README.md#deleterelation) - Delete a relation
* [GetRelation](docs/sdks/relations/README.md#getrelation) - Get a relation
* [ListRelations](docs/sdks/relations/README.md#listrelations) - List relations
* [UpdateRelation](docs/sdks/relations/README.md#updaterelation) - Update a relation

### [Tasks](docs/sdks/tasks/README.md)

* [CreateTask](docs/sdks/tasks/README.md#createtask) - Create a task
* [DeleteTask](docs/sdks/tasks/README.md#deletetask) - Delete a task
* [GetTask](docs/sdks/tasks/README.md#gettask) - Get a task
* [ListTasks](docs/sdks/tasks/README.md#listtasks) - List tasks
* [UpdateTask](docs/sdks/tasks/README.md#updatetask) - Update a task

### [Users](docs/sdks/users/README.md)

* [CurrentUser](docs/sdks/users/README.md#currentuser) - Current User
* [ListUsers](docs/sdks/users/README.md#listusers) - List users
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Error Handling -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                                         | Status Code                                          | Content Type                                         |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| sdkerrors.ListContractTypesResponseBody              | 401                                                  | application/json                                     |
| sdkerrors.ListContractTypesContractTypesResponseBody | 403                                                  | application/json                                     |
| sdkerrors.SDKError                                   | 400-600                                              | */*                                                  |

### Example

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
	s := contractifyproduction.New()

	ctx := context.Background()
	res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
		Company: 839467,
	})
	if err != nil {

		var e *sdkerrors.ListContractTypesResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.ListContractTypesContractTypesResponseBody
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->

<!-- Start Server Selection -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://app.contractify.be` | None |

#### Example

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
		contractifyproduction.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		contractifyproduction.WithServerURL("https://app.contractify.be"),
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
<!-- End Server Selection -->

<!-- Start Custom HTTP Client -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->



<!-- Start Authentication -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                  | Type                  | Scheme                |
| --------------------- | --------------------- | --------------------- |
| `OAuth2`              | oauth2                | OAuth2 token          |
| `PersonalAccessToken` | http                  | HTTP Bearer           |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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
<!-- End Authentication -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
