// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"ContractifyProduction/pkg/models/shared"
	"net/http"
)

type CreateTaskSecurity struct {
	OAuth2              string `security:"scheme,type=oauth2,name=Authorization"`
	PersonalAccessToken string `security:"scheme,type=http,subtype=bearer,name=Authorization"`
}

type CreateTaskRequest struct {
	TaskWrite *shared.TaskWrite `request:"mediaType=application/json"`
	// Id of the company
	Company int64 `pathParam:"style=simple,explode=false,name=company"`
}

type CreateTask422ApplicationJSONErrors struct {
	Errors []string `json:"errors,omitempty"`
	Field  *string  `json:"field,omitempty"`
}

// CreateTask422ApplicationJSON - Invalid data posted
type CreateTask422ApplicationJSON struct {
	Errors  []CreateTask422ApplicationJSONErrors `json:"errors,omitempty"`
	Message *string                              `json:"message,omitempty"`
}

// CreateTask403ApplicationJSON - Forbidden
type CreateTask403ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// CreateTask401ApplicationJSON - Unauthenticated
type CreateTask401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

// CreateTask200ApplicationJSON - OK
type CreateTask200ApplicationJSON struct {
	Data *shared.TaskRead `json:"data,omitempty"`
}

type CreateTaskResponse struct {
	ContentType string
	StatusCode  int
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
