# ContractifyProduction

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/contractify-api-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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
    operationSecurity := operations.ListContractTypesSecurity{
            OAuth2: "",
            PersonalAccessToken: "",
        }

    ctx := context.Background()
    res, err := s.ContractTypes.ListContractTypes(ctx, operations.ListContractTypesRequest{
        Company: 548814,
    }, operationSecurity)
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

### [Subfolders](docs/sdks/subfolders/README.md)

* [ListSubfolders](docs/sdks/subfolders/README.md#listsubfolders) - List subfolders

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

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
