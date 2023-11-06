// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type CreateTaskRequest struct {
	TaskWrite *shared.TaskWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

func (o *CreateTaskRequest) GetTaskWrite() *shared.TaskWrite {
	if o == nil {
		return nil
	}
	return o.TaskWrite
}

func (o *CreateTaskRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

type CreateTask422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *CreateTask422ApplicationJSONErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *CreateTask422ApplicationJSONErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// CreateTask422ApplicationJSON - Invalid data posted
type CreateTask422ApplicationJSON struct {
	Errors  []CreateTask422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                              `json:"message,omitempty"`
}

func (o *CreateTask422ApplicationJSON) GetErrors() []CreateTask422ApplicationJSONErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *CreateTask422ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateTask403ApplicationJSON - Forbidden
type CreateTask403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CreateTask403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateTask401ApplicationJSON - Unauthenticated
type CreateTask401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *CreateTask401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// CreateTask200ApplicationJSON - OK
type CreateTask200ApplicationJSON struct {
	Data *shared.TaskRead `json:"data,omitempty"`
}

func (o *CreateTask200ApplicationJSON) GetData() *shared.TaskRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type CreateTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	CreateTask200ApplicationJSONObject *CreateTask200ApplicationJSON
	// Unauthenticated
	CreateTask401ApplicationJSONObject *CreateTask401ApplicationJSON
	// Forbidden
	CreateTask403ApplicationJSONObject *CreateTask403ApplicationJSON
	// Invalid data posted
	CreateTask422ApplicationJSONObject *CreateTask422ApplicationJSON
}

func (o *CreateTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTaskResponse) GetCreateTask200ApplicationJSONObject() *CreateTask200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateTask200ApplicationJSONObject
}

func (o *CreateTaskResponse) GetCreateTask401ApplicationJSONObject() *CreateTask401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateTask401ApplicationJSONObject
}

func (o *CreateTaskResponse) GetCreateTask403ApplicationJSONObject() *CreateTask403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateTask403ApplicationJSONObject
}

func (o *CreateTaskResponse) GetCreateTask422ApplicationJSONObject() *CreateTask422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.CreateTask422ApplicationJSONObject
}
