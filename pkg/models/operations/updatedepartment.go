// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type UpdateDepartmentRequest struct {
	DepartmentWrite *shared.DepartmentWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the department
	Department int64 `pathParam:"style=simple,explode=false,name=department"`
}

func (o *UpdateDepartmentRequest) GetDepartmentWrite() *shared.DepartmentWrite {
	if o == nil {
		return nil
	}
	return o.DepartmentWrite
}

func (o *UpdateDepartmentRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *UpdateDepartmentRequest) GetDepartment() int64 {
	if o == nil {
		return 0
	}
	return o.Department
}

type UpdateDepartment422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *UpdateDepartment422ApplicationJSONErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateDepartment422ApplicationJSONErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// UpdateDepartment422ApplicationJSON - Invalid data posted
type UpdateDepartment422ApplicationJSON struct {
	Errors  []UpdateDepartment422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                                    `json:"message,omitempty"`
}

func (o *UpdateDepartment422ApplicationJSON) GetErrors() []UpdateDepartment422ApplicationJSONErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateDepartment422ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDepartment404ApplicationJSON - Not Found
type UpdateDepartment404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDepartment404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDepartment403ApplicationJSON - Forbidden
type UpdateDepartment403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDepartment403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDepartment401ApplicationJSON - Unauthenticated
type UpdateDepartment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateDepartment401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateDepartment200ApplicationJSON - OK
type UpdateDepartment200ApplicationJSON struct {
	Data *shared.DepartmentRead `json:"data,omitempty"`
}

func (o *UpdateDepartment200ApplicationJSON) GetData() *shared.DepartmentRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateDepartmentResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdateDepartment200ApplicationJSONObject *UpdateDepartment200ApplicationJSON
	// Unauthenticated
	UpdateDepartment401ApplicationJSONObject *UpdateDepartment401ApplicationJSON
	// Forbidden
	UpdateDepartment403ApplicationJSONObject *UpdateDepartment403ApplicationJSON
	// Not Found
	UpdateDepartment404ApplicationJSONObject *UpdateDepartment404ApplicationJSON
	// Invalid data posted
	UpdateDepartment422ApplicationJSONObject *UpdateDepartment422ApplicationJSON
}

func (o *UpdateDepartmentResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateDepartmentResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateDepartmentResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateDepartmentResponse) GetUpdateDepartment200ApplicationJSONObject() *UpdateDepartment200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDepartment200ApplicationJSONObject
}

func (o *UpdateDepartmentResponse) GetUpdateDepartment401ApplicationJSONObject() *UpdateDepartment401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDepartment401ApplicationJSONObject
}

func (o *UpdateDepartmentResponse) GetUpdateDepartment403ApplicationJSONObject() *UpdateDepartment403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDepartment403ApplicationJSONObject
}

func (o *UpdateDepartmentResponse) GetUpdateDepartment404ApplicationJSONObject() *UpdateDepartment404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDepartment404ApplicationJSONObject
}

func (o *UpdateDepartmentResponse) GetUpdateDepartment422ApplicationJSONObject() *UpdateDepartment422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateDepartment422ApplicationJSONObject
}
