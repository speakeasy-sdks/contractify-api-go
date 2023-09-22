// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type UpdateTaskRequest struct {
	TaskUpdate *shared.TaskUpdate `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the task
	Task int64 `pathParam:"style=simple,explode=false,name=task"`
}

func (o *UpdateTaskRequest) GetTaskUpdate() *shared.TaskUpdate {
	if o == nil {
		return nil
	}
	return o.TaskUpdate
}

func (o *UpdateTaskRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *UpdateTaskRequest) GetTask() int64 {
	if o == nil {
		return 0
	}
	return o.Task
}

type UpdateTask422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

func (o *UpdateTask422ApplicationJSONErrors) GetErrors() []string {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateTask422ApplicationJSONErrors) GetField() *string {
	if o == nil {
		return nil
	}
	return o.Field
}

// UpdateTask422ApplicationJSON - Invalid data posted
type UpdateTask422ApplicationJSON struct {
	Errors  []UpdateTask422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                              `json:"message,omitempty"`
}

func (o *UpdateTask422ApplicationJSON) GetErrors() []UpdateTask422ApplicationJSONErrors {
	if o == nil {
		return nil
	}
	return o.Errors
}

func (o *UpdateTask422ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateTask404ApplicationJSON - Not Found
type UpdateTask404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateTask404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateTask403ApplicationJSON - Forbidden
type UpdateTask403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateTask403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateTask401ApplicationJSON - Unauthenticated
type UpdateTask401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *UpdateTask401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// UpdateTask200ApplicationJSON - OK
type UpdateTask200ApplicationJSON struct {
	Data *shared.TaskRead `json:"data,omitempty"`
}

func (o *UpdateTask200ApplicationJSON) GetData() *shared.TaskRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type UpdateTaskResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// OK
	UpdateTask200ApplicationJSONObject *UpdateTask200ApplicationJSON
	// Unauthenticated
	UpdateTask401ApplicationJSONObject *UpdateTask401ApplicationJSON
	// Forbidden
	UpdateTask403ApplicationJSONObject *UpdateTask403ApplicationJSON
	// Not Found
	UpdateTask404ApplicationJSONObject *UpdateTask404ApplicationJSON
	// Invalid data posted
	UpdateTask422ApplicationJSONObject *UpdateTask422ApplicationJSON
}

func (o *UpdateTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateTaskResponse) GetUpdateTask200ApplicationJSONObject() *UpdateTask200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTask200ApplicationJSONObject
}

func (o *UpdateTaskResponse) GetUpdateTask401ApplicationJSONObject() *UpdateTask401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTask401ApplicationJSONObject
}

func (o *UpdateTaskResponse) GetUpdateTask403ApplicationJSONObject() *UpdateTask403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTask403ApplicationJSONObject
}

func (o *UpdateTaskResponse) GetUpdateTask404ApplicationJSONObject() *UpdateTask404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTask404ApplicationJSONObject
}

func (o *UpdateTaskResponse) GetUpdateTask422ApplicationJSONObject() *UpdateTask422ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.UpdateTask422ApplicationJSONObject
}
