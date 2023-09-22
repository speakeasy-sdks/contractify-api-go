// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type CreateDepartmentRequest struct {
	DepartmentWrite *shared.DepartmentWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

func (o *CreateDepartmentRequest) GetDepartmentWrite() *shared.DepartmentWrite {
	if o == nil {
		return nil
	}
	return o.DepartmentWrite
}

func (o *CreateDepartmentRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

type CreateDepartment422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *CreateDepartment422ApplicationJSONErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *CreateDepartment422ApplicationJSONErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// CreateDepartment422ApplicationJSON - Invalid data posted
type CreateDepartment422ApplicationJSON struct {
	Errors  []CreateDepartment422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                                    `json:"message,omitempty"`
}

func (o *CreateDepartment422ApplicationJSON) GetErrors() []CreateDepartment422ApplicationJSONErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *CreateDepartment422ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateDepartment403ApplicationJSON - Forbidden
type CreateDepartment403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CreateDepartment403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateDepartment401ApplicationJSON - Unauthenticated
type CreateDepartment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CreateDepartment401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateDepartment201ApplicationJSON - Created
type CreateDepartment201ApplicationJSON struct {
	Data *shared.DepartmentRead `json:"data,omitempty"`
}

func (o *CreateDepartment201ApplicationJSON) GetData() *shared.DepartmentRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateDepartmentResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Created
	CreateDepartment201ApplicationJSONObject *CreateDepartment201ApplicationJSON
	// Unauthenticated
	CreateDepartment401ApplicationJSONObject *CreateDepartment401ApplicationJSON
	// Forbidden
	CreateDepartment403ApplicationJSONObject *CreateDepartment403ApplicationJSON
	// Invalid data posted
	CreateDepartment422ApplicationJSONObject *CreateDepartment422ApplicationJSON
}

func (o *CreateDepartmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateDepartmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateDepartmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateDepartmentResponse) GetCreateDepartment201ApplicationJSONObject() *CreateDepartment201ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateDepartment201ApplicationJSONObject
}

func (o *CreateDepartmentResponse) GetCreateDepartment401ApplicationJSONObject() *CreateDepartment401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateDepartment401ApplicationJSONObject
}

func (o *CreateDepartmentResponse) GetCreateDepartment403ApplicationJSONObject() *CreateDepartment403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateDepartment403ApplicationJSONObject
}

func (o *CreateDepartmentResponse) GetCreateDepartment422ApplicationJSONObject() *CreateDepartment422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateDepartment422ApplicationJSONObject
}
