// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type GetTaskRequest struct {
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
	// Id of the task
	Task int64 `pathParam:"style=simple,explode=false,name=task"`
}

func (o *GetTaskRequest) GetCompany() int64 {
	if o == nil {
		return 0
	}
	return o.Company
}

func (o *GetTaskRequest) GetTask() int64 {
	if o == nil {
		return 0
	}
	return o.Task
}

// GetTask404ApplicationJSON - Not Found
type GetTask404ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetTask404ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetTask403ApplicationJSON - Forbidden
type GetTask403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetTask403ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetTask401ApplicationJSON - Unauthenticated
type GetTask401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

func (o *GetTask401ApplicationJSON) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

// GetTask200ApplicationJSON - OK
type GetTask200ApplicationJSON struct {
	Data *shared.TaskRead `json:"data,omitempty"`
}

func (o *GetTask200ApplicationJSON) GetData() *shared.TaskRead {
	if o == nil {
		return nil
	}
	return o.Data
}

type GetTaskResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	GetTask200ApplicationJSONObject *GetTask200ApplicationJSON
	// Unauthenticated
	GetTask401ApplicationJSONObject *GetTask401ApplicationJSON
	// Forbidden
	GetTask403ApplicationJSONObject *GetTask403ApplicationJSON
	// Not Found
	GetTask404ApplicationJSONObject *GetTask404ApplicationJSON
}

func (o *GetTaskResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTaskResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTaskResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTaskResponse) GetGetTask200ApplicationJSONObject() *GetTask200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTask200ApplicationJSONObject
}

func (o *GetTaskResponse) GetGetTask401ApplicationJSONObject() *GetTask401ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTask401ApplicationJSONObject
}

func (o *GetTaskResponse) GetGetTask403ApplicationJSONObject() *GetTask403ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTask403ApplicationJSONObject
}

func (o *GetTaskResponse) GetGetTask404ApplicationJSONObject() *GetTask404ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetTask404ApplicationJSONObject
}
